package models

type ExpenseTypes struct {
	Type        string `json:"expense_type"`
	Description string `json:"description"`
	Color       string `json:"color"`
}