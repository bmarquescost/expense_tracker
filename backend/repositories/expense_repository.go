package repositories

import (
	"log"
    "database/sql"
    "github.com/bmarquescost/expense-tracker/models"
)

type ExpenseRepository struct {
    DB *sql.DB
}

type ExpenseWithType struct {
	UserID       		   string 
	ID           		   string 
	Description  		   string 
	Amount       		   float64 
	Date	     		   string 
    ExpenseType  		   string 
	CronSchedule 		   string  
	StartDate	 		   string 
    EndDate 	 		   string 
	ExpenseTypeDescription string
	ExpenseTypeColor       string
}


func (r *ExpenseRepository) ExpenseTypeAlreadyExists(expenseType string) bool {
	query := "SELECT COUNT(*) FROM expense_types WHERE expense_type = ?"
	rows, err := r.DB.Query(query, expenseType)
	if err != nil {
		log.Fatal("%s", err)
		return false
	}
	defer rows.Close() 
	
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	log.Printf("Count: %d", count)

	return count != 0
} 

func (r *ExpenseRepository) UpsertExpenseType(expenseType *models.ExpenseTypes)  error {
    query := `
		INSERT INTO expense_types (expense_type, title, color) 
		VALUES (?,?,?)
		ON DUPLICATE KEY UPDATE
			expense_type = VALUES(expense_type),
			title = VALUES(title),
			color = VALUES(color)
	`
	
	_, err := r.DB.Exec(query, expenseType.Type, expenseType.Description, expenseType.Color)
	log.Printf("Query: %s", query)
	log.Printf("%s", err)
	return err
}

func (r *ExpenseRepository) UpsertExpense(expense *models.Expense) error {
	exists := r.ExpenseTypeAlreadyExists(expense.ExpenseType)
	log.Printf("%d", exists)
	if !exists {
		newExpenseTypes := models.ExpenseTypes{expense.ExpenseType, "", ""}
		r.UpsertExpenseType(&newExpenseTypes)
	}

    query := `
		INSERT INTO expenses (username, expense_title, title, amount, date, expense_type, cron_schedule, start_date, end_date) 
		VALUES (?,?,?,?,?,?,?,?,?)
		ON DUPLICATE KEY UPDATE
			username = VALUES(username),
			expense_title = VALUES(expense_title),
			title = VALUES(title),
			amount = VALUES(amount),
			date = VALUES(date),
			expense_type = VALUES(expense_type),
			cron_schedule = VALUES(cron_schedule),
			start_date = VALUES(start_date),
			end_date = VALUES(end_date)
		`
	_, err := r.DB.Exec(query, expense.UserID, expense.ID, expense.Description, expense.Amount, expense.Date, expense.ExpenseType, expense.CronSchedule, expense.StartDate, expense.EndDate)

    return err
}

func (r *ExpenseRepository) GetUserExpenses(user string) ([]ExpenseWithType, error) {
	query := `
		SELECT 
			expenses.username,
			expenses.expense_title,
			expenses.title,
			expenses.amount,
			expenses.date,
			expenses.expense_type,
			expenses.cron_schedule,
			expenses.start_date,
			expenses.end_date,
			expense_types.title,
			expense_types.color
		FROM expenses
		JOIN expense_types 
		USING (expense_type) 
		WHERE username = ? 
	`
	rows, err := r.DB.Query(query, user)
	if err != nil {
		return nil, err
	}

	var userExpenses []ExpenseWithType
	for rows.Next() {
		var userExpense ExpenseWithType

		err := rows.Scan(
			&userExpense.UserID,
			&userExpense.ID,
			&userExpense.Description,
			&userExpense.Amount,
			&userExpense.Date,
			&userExpense.ExpenseType,
			&userExpense.CronSchedule,
			&userExpense.StartDate,
			&userExpense.EndDate,
			&userExpense.ExpenseTypeDescription,
			&userExpense.ExpenseTypeColor,
		)

		if err != nil {
			return nil, err 
		}

		userExpenses = append(userExpenses, userExpense)
	}

	if err := rows.Err(); err != nil {
		return nil, err 
	}

	return userExpenses, nil
}

func (r *ExpenseRepository) DeleteExpense(userID string, expenseTitle string) error {
	query := `DELETE FROM expenses WHERE username = ? AND expense_title = ?`
	_, err := r.DB.Exec(query, userID, expenseTitle)
	return err
}

func (r *ExpenseRepository) DeleteExpenseType(expenseType models.ExpenseTypes) error {
	query := `DELETE FROM expense_types WHERE expense_type = ?`
	_, err := r.DB.Exec(query, expenseType.Type)
	return err
}