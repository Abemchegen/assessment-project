package main

import (
	"log"

	"loantracker/deliverables/controller"
	"loantracker/deliverables/router"
	"loantracker/infrastructure"
	"loantracker/repo"
	"loantracker/usecase"
)

func main() {
	// Get the database connection from the infrastructure package
	db, err := infrastructure.InitializeMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// // Initialize the repository
	// loanRepository := repo.NewLoanRepository(db)

	// // Initialize the use case
	// loanUsecase := usecase.NewLoanUsecase(loanRepository)

	// // Initialize the controller
	// loanController := controller.NewLoanController(loanUsecase)

	// Initialize the user repository
	userRepository := repo.NewUserRepository(db)

	// Initialize the user use case
	userUsecase := usecase.NewUserUsecase(userRepository)

	// Initialize the user controller
	userController := controller.NewUserController(userUsecase)

	// Initialize the router
	r := router.InitializeRouter(userController)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
