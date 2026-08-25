package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sawsdev/expense-tracker/internal/file"
)

var startExpensesCmd = &cobra.Command{
	Use: "init",
	Short: "Create/verify the necessary files for the expense list program",
	Long: "Create assets/db/expenses.csv file for storing the expense list if it doesn't exists",
	Run: createExpenseList,
}

func init() {
	RootCmd.AddCommand(startExpensesCmd)
}


func createExpenseList(cmd * cobra.Command, args [] string) {
	
	if !file.FileExists("expenses.csv") {
		file.MakeDirectory("assets")
		file.MakeDirectory("assets/db")
		file.CreateFile("expenses.csv")
	} else {
		fmt.Println("Expense file found, starting.")
		file.ReadCSVFile("expenses.csv")
	}

}