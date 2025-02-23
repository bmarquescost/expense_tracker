package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/bmarquescost/expense-tracker/models"
	"github.com/bmarquescost/expense-tracker/repositories"
)

type ExpenseHandler struct {
	Repo *repositories.ExpenseRepository
}

type DeleteExpenseRequest struct {
	UserID       string `json:"username"`
	ExpenseTitle string `json:"expense_title"`
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	log.Printf("Inserting expense: ", expense)
	if err := h.Repo.UpsertExpense(&expense); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to insert expense: ", expense)
        return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) CreateExpenseType(w http.ResponseWriter, r *http.Request) {
	var expenseType models.ExpenseTypes
	err := json.NewDecoder(r.Body).Decode(&expenseType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	log.Printf("Inserting expense: ", expenseType)
	if err := h.Repo.UpsertExpenseType(&expenseType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to insert expense type: ", expenseType)
        return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expenseType)
}

func (h *ExpenseHandler) GetUserExpenses(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	log.Printf("Retrieving all expenses from user: ", user)

	if user != "" {
		userExpenses, err := h.Repo.GetUserExpenses(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Failed to get user %s expenses", user)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userExpenses)
		return 
	}

	http.Error(w, "Empty user not accepted. Please provide a valid user", http.StatusInternalServerError)
	log.Println("Empty user not accepted. Please provide a valid user")
}

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	var expenseToDelete DeleteExpenseRequest
	err := json.NewDecoder(r.Body).Decode(&expenseToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	log.Printf("Deleting expense relative to user %s and title %s", expenseToDelete.UserID, expenseToDelete.ExpenseTitle)
	if err := h.Repo.DeleteExpense(expenseToDelete.UserID, expenseToDelete.ExpenseTitle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to delete expense relative to user %s and title %s", expenseToDelete.UserID, expenseToDelete.ExpenseTitle)
        return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expenseToDelete)
}

func (h *ExpenseHandler) DeleteExpenseType(w http.ResponseWriter, r *http.Request) {
	var expenseType models.ExpenseTypes
	err := json.NewDecoder(r.Body).Decode(&expenseType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	log.Printf("Deleting expense type: ", expenseType)
	if err := h.Repo.DeleteExpenseType(expenseType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to delete expense type: ", expenseType)
        return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expenseType)
}