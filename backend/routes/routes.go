package routes

import (
    "github.com/gorilla/mux"
    "github.com/bmarquescost/expense-tracker/handlers"
    "github.com/bmarquescost/expense-tracker/repositories"
    "database/sql"
)

func SetupRouter(db *sql.DB) *mux.Router {
    router := mux.NewRouter()
    
	userRepo := &repositories.UserRepository{DB: db}
    expenseRepo := &repositories.ExpenseRepository{DB: db}

    userHandler := &handlers.UserHandler{Repo: userRepo}
    expenseHandler := &handlers.ExpenseHandler{Repo: expenseRepo}

	// User Handlers
    router.HandleFunc("/register-user", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login-user", userHandler.AuthenticateUser).Methods("POST")

	// Expense Handlers
    router.HandleFunc("/upsert-expense", expenseHandler.CreateExpense).Methods("POST")
	router.HandleFunc("/upsert-expense-type", expenseHandler.CreateExpenseType).Methods("POST")
	router.HandleFunc("/get-user-expenses", expenseHandler.GetUserExpenses).Methods("GET")
	router.HandleFunc("/delete-expense", expenseHandler.DeleteExpense).Methods("POST")
	router.HandleFunc("/delete-expense-type", expenseHandler.DeleteExpenseType).Methods("POST")

    return router
}