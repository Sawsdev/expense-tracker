package expenselist

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

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


}

func ShowExpenses(expenseList *ExpenseList) {
   
	var writer strings.Builder
	writer.WriteString(expenseHeader)
	for _, expense := range expenseList.Expenses {
		//CHANGE: Adjusted the way to show strings in a efficient way
		formattedString := fmt.Sprintf("%d\t%s\t%s\t\t\t%d\n", expense.Id, expense.Date, expense.Description, expense.Amount)
		writer.WriteString(formattedString)
		
	}
	fmt.Print(writer.String())
	writer.Reset()
}

func GetSingleExpense(expenseList *ExpenseList, id int){
	obtainedExpense, _ := getExpense(expenseList, id)
	showExpense(obtainedExpense)
}

func showExpense(loggedExpense expense.Expense){
	formattedString := fmt.Sprintf("%d\t%s\t%s\t\t\t%d\n", loggedExpense.Id, loggedExpense.Date, loggedExpense.Description, loggedExpense.Amount)
	fmt.Print(expenseHeader+"\n"+formattedString)
}

func getExpense(expenseList *ExpenseList, id int) (expense.Expense, int) {

	defaultExpense := expense.NewExpense(0, "", "", 0)
	index := 0
	for i, expense := range expenseList.Expenses {
		if id == expense.Id {
			defaultExpense = expense
			index = i
		}
	}
	return defaultExpense, index
}

func DeleteExpense(expenseList *ExpenseList, id int){
	_, index := getExpense(expenseList, id)
	if index == 0 {
		fmt.Println("Expense not found")
		return 
	}
	expenseList.Expenses = slices.DeleteFunc(expenseList.Expenses, func(e expense.Expense)bool {
		return e.Id == id
	})
	fmt.Println("Expense Deleted")

}

func UpdateExpense(expenseList *ExpenseList, id int, description string, date string, amount int){

	expenseToUpdate, index := getExpense(expenseList, id)
	if index == 0 {
		fmt.Println("Expense not found")
		return
	}

	expenseToUpdate.Description = description
	expenseToUpdate.Date = date
	expenseToUpdate.Amount = amount

	expenseList.Expenses[index] = expenseToUpdate
	fmt.Println("Expense updated succesfully")

}

func SummarizeExpenses(expenseList *ExpenseList, day int, month int, year int) (int, string){
	total := 0
	currentMonth := ""
	stringGivenDate := preformatStringDate(day, month, year)
	givenDate, err := time.Parse(dateLayout, stringGivenDate)
	if err != nil {
		log.Fatal("Error parsing given date", err)
	}
	currentMonth = givenDate.Month().String()
	for _, expense := range expenseList.Expenses {
		if(month > 0 && year > 0){
			actualDate, err := time.Parse(dateLayout, expense.Date)
			if err != nil {
				log.Fatal("Error parsing date")
				return 0, ""
			}
			if((day>0 && day == actualDate.Day()) && month == int(actualDate.Month()) && year == actualDate.Year()){
				total = total + expense.Amount

			}else if (month == int(actualDate.Month()) && year == actualDate.Year()){
				total = total + expense.Amount

			} else if (day <= 0 && month <= 0 && actualDate.Year() == year) {
				total = total + expense.Amount
			}
		} else {
			total = total + expense.Amount
		}
	}
	return total, currentMonth
}

func preformatStringDate (day int, month int, year int) string{
	stringDay := ""
	stringMonth := ""
	stringYear := ""
	if day <= 0  || day > 31{
		stringDay = "01"
	}else {
		stringDay = fmt.Sprintf("%d", day)
	}
	if month <= 0 || month > 12 {
		stringMonth = "01"
	} else if(month < 9){
		stringMonth = fmt.Sprintf("0%d", month)
	}else {
		stringMonth = fmt.Sprintf("%d", month)
	}
	if year < 1000 {
		stringYear = "1000"
	} else {
		stringYear = fmt.Sprintf("%d", year)
	}
	return fmt.Sprintf("%s-%s-%s", stringYear, stringMonth, stringDay)
}
