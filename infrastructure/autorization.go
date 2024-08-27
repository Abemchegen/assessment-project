package infrastructure

import (
	"context"
	"loantracker/domain"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JWTMiddleware is a middleware that extracts JWT claims and sets them in the context
func JWTMiddleware(c *gin.Context) {
	// Get the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	// Extract the token from the "Bearer <token>" format
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	// Parse the token using the ParseJWT function
	claims, err := ParseJWT(tokenString)
	if err != nil {

		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {

			db, err := InitializeMongoDB()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
				c.Abort()
			}
			defer db.Disconnect(context.Background())
			collection := db.Database("Loan-tracker").Collection("Users")
			var user domain.User
			userid, err := primitive.ObjectIDFromHex(claims.Id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
				c.Abort()
			}
			err = collection.FindOne(context.TODO(), bson.M{"_id": userid}).Decode(&user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
				c.Abort()
			}
			if user.RefreshToken == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
				c.Abort()
			}

			// Validate the refresh token
			refreshClaims, err := ParseJWT(user.RefreshToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
				c.Abort()
				return
			}

			// Generate a new access token
			newToken, err := GenerateJWT(refreshClaims.Name, refreshClaims.ID, refreshClaims.Role, true)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
				c.Abort()
				return
			}

			// Set the new access token in the response header
			c.Header("Authorization", "Bearer "+newToken)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}

	// Set the claims in the context
	c.Set("claims", claims)
	c.Next()
}
