package models

type Expense struct {
	UserID       string  `json:"username"`
	ID           string  `json:"expense_title"`
	Description  string  `json:"description"`
	Amount       float64 `json:"amount"` 
	Date	     string  `json:"date"`
    ExpenseType  string  `json:"expense_type"`
	CronSchedule string  `json:"cron_schedule"` 
	StartDate	 string  `json:"start_date"`
    EndDate 	 string  `json:"end_date"`
} 