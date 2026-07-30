package controller

import (
	"fmt"
	"log"
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
	id := convertFileToExpenseList()
	expenselist.AddNewExpense(&expenses,description,amount, id)
	dataForWriting := convertExpenseListToFile()
	file.WriteCSVFile("expenses.csv", dataForWriting)
	ShowExpenseList()
}

func ShowExpenseList() {
	fmt.Println(expenses.Expenses)
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
		expenselist.AddNewExpense(&expenses, description, amount, actualId)
	}
	fmt.Println(fileData)
	return  lastId
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