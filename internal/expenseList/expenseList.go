package expenselist

import (
	"fmt"
	"time"
	"strings"

	"github.com/sawsdev/expense-tracker/internal/expense"
)

const (
	dateLayout = "2006-01-02"
	expenseHeader = "ID\tDate\t\tDescription\t\tAmount\n"

)

type ExpenseList struct {
	Expenses []expense.Expense
}

func NewExpenseList() ExpenseList {
	return ExpenseList{
		Expenses: []expense.Expense{},
	}
}

func AddNewExpense(expenseList *ExpenseList, description string, amount int, id int, date string) {
	now := time.Now().Local().UTC()
	if date == "" {
		date = now.Format(dateLayout)
	}
	newExpense := expense.NewExpense(
		id, //added custom id to keep it consistent for the file
		description,
		date,
		amount)
	expenseList.Expenses = append(expenseList.Expenses, newExpense)
	fmt.Println("New Expense added")

}

func ShowExpenses(expenseList *ExpenseList) {
   
	var writer strings.Builder
	writer.WriteString(expenseHeader)
	for _, expense := range expenseList.Expenses {
		//CHANGE: Adjusted the way to show strings in a efficient way
		writer.WriteString(fmt.Sprintf("%d\t%s\t%s\t\t\t%d\n", expense.Id, expense.Date, expense.Description, expense.Amount))
		
	}
	fmt.Print(writer.String())
}

func GetSingleExpense(expenseList *ExpenseList, id int){
	obtainedExpense := getExpense(expenseList, id)
	showExpense(obtainedExpense)
}

func showExpense(loggedExpense expense.Expense){
	formattedString := fmt.Sprintf("%d\t%s\t%s\t\t\t%d\n", loggedExpense.Id, loggedExpense.Date, loggedExpense.Description, loggedExpense.Amount)
	fmt.Print(expenseHeader+"\n"+formattedString)
}

func getExpense(expenseList *ExpenseList, id int) expense.Expense {

	defaultExpense := expense.NewExpense(0, "", "", 0)
	for _, expense := range expenseList.Expenses {
		if id == expense.Id {
			defaultExpense = expense
		}
	}
	return defaultExpense
}
