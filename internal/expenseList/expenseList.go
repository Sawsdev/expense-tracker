package expenselist

import (
	"fmt"
	"time"

	"github.com/sawsdev/expense-tracker/internal/expense"
)

const (
	dateLayout = "2006-01-02"
)

type ExpenseList struct {
	Expenses []expense.Expense
}

func NewExpenseList() ExpenseList {
	return ExpenseList{
		Expenses: []expense.Expense{},
	}
}

func AddNewExpense(expenseList *ExpenseList, description string, amount int) {
	now := time.Now().Local().UTC()
	newExpense := expense.NewExpense(
		len(expenseList.Expenses),
		description,
		now.Format(dateLayout),
		amount)
	expenseList.Expenses = append(expenseList.Expenses, newExpense)
	fmt.Println("New Expense added")
	
}