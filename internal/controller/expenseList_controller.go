package controller

import (
	"fmt"

	"github.com/sawsdev/expense-tracker/internal/expenseList"
)
var expenses = expenselist.ExpenseList{}

func GetExpenseList()  expenselist.ExpenseList{
	return expenses
}

func CreateExpenseList() expenselist.ExpenseList{
	expenses = expenselist.NewExpenseList()
	return expenses
}

func AddNewExpenseToList(description string, amount int) {
	expenselist.AddNewExpense(&expenses,description,amount)
	ShowExpenseList()
}

func ShowExpenseList() {
	fmt.Println(expenses.Expenses)
}