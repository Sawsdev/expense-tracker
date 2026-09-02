package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/sawsdev/expense-tracker/internal/expenseList"
	"github.com/sawsdev/expense-tracker/internal/file"
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
	if(!isValidExpenseInput(description,amount)){
		return
	}
	id := convertFileToExpenseList()
	expenselist.AddNewExpense(&expenses,description,amount, id, "")
	saveExpenseListInCSV()
	fmt.Println("New Expense added")
}

func ShowExpenseList() {
	//Its called to fill the expense list before showing.
	convertFileToExpenseList()
	expenselist.ShowExpenses(&expenses)
}

func ShowSingleExpense(id int){
	convertFileToExpenseList()
	expenselist.GetSingleExpense(&expenses, id)
}

func UpdateExpense(id int, description string, date string, amount int){
	if(!isValidExpenseInput(description,amount)){
		return
	}
	if(!isValidDateInput(date)){
		return
	}
	convertFileToExpenseList()
	expenselist.UpdateExpense(&expenses, id, description, date, amount)
	saveExpenseListInCSV()
}

func DeleteExpense(id int){
	convertFileToExpenseList()
	expenselist.DeleteExpense(&expenses, id)
	saveExpenseListInCSV()
}

func GetExpenseSummary(day int, month int, year int){
	if(day > 0 && (month <= 0 || year <= 0)){
		fmt.Println("month and year are required")
		return
	}else if(month > 0 && year <= 0){
		fmt.Println("Year is required to get the summary")
		return
	}
	convertFileToExpenseList()
	summary, currentMonth := expenselist.SummarizeExpenses(&expenses,day, month, year)
	if month > 0 && year > 0 {
		if (day > 0) {
		fmt.Printf("Total expenses for the %d of %s of year %d: %d\n", day, currentMonth, year, summary)
		}
		fmt.Printf("Total expenses for %s: %d\n", currentMonth, summary)
	}else if(month <= 0 && year > 0){
		fmt.Printf("Total expenses for year %d: %d\n", year, summary)
	} else {
		fmt.Printf("Total expenses: %d\n", summary)
	}

}

func convertFileToExpenseList() int {
	fileData := file.ReadCSVFile("expenses.csv")
	lastId , err := strconv.Atoi(fileData[len(fileData)-1][0])
	if err != nil {
		log.Fatal("Error converting last id: ", err)
	}
	for i, expense := range fileData {
		if i == 0 { //skip header
			continue
		}
		actualId, err := strconv.Atoi(expense[0])
		if err != nil {
			log.Fatal("Error parsing id: ", err)
		}
		amount, err := strconv.Atoi(expense[3])
		if err != nil {
			log.Fatal("Error parsing amount: ", err)
		}
		description := expense[2]
		expenselist.AddNewExpense(&expenses, description, amount, actualId, expense[1])
	}
	return  lastId + 1
}

func convertExpenseListToFile() [][] string {
	convertedList :=  [][] string {
		{"ID","Date","Description","Amount"},
	}
	for _, expense := range expenses.Expenses {
		id := strconv.Itoa(expense.Id)
		amount := strconv.Itoa(expense.Amount) 
		row := []string {id, expense.Date, expense.Description, amount}
		convertedList = append(convertedList, row)
	}

	return convertedList
}

func saveExpenseListInCSV(){
	dataForWriting := convertExpenseListToFile()
	file.WriteCSVFile("expenses.csv", dataForWriting)
}

func isValidExpenseInput(description string, amount int)bool {

	if(len(description) > 80){
		fmt.Println("Description exceeded 80 characters")
		return false
	}
	if(amount < 0){
		fmt.Println("Amount cant be less than 0")
		return false
	}
	return true
}

func isValidDateInput(date string) bool{ 

	rule, error := regexp.Compile("[0-9]{4}-[0-9]{2}-[0-9]{1,31}")
	if error != nil {
		log.Fatal("Error compiling the regular expression")
	}

	if(!rule.MatchString(date)){
		fmt.Println("Incorrect date format. expected: YYYY-MM-DD")
		return false
	}
	return true
}