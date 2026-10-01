package cmd

/**
* command form
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
**/

import (
	//"fmt"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	"github.com/sawsdev/expense-tracker/internal/controller"
)

var (
	//flags
	description string
	category    string
	amount      int
	//command
	addExpenseCmd = &cobra.Command{
		Use:   "add",
		Short: "Add new expense entry",
		Long:  "Add a new expense entry to the list including --description and --amount",
		Run:   addExpense,
	}
)

func init() {
	addExpenseCmd.Flags().StringVarP(&description, "description", "d", "", "expense description")
	addExpenseCmd.Flags().StringVarP(&category, "category", "c", "", "expense category")
	addExpenseCmd.Flags().IntVarP(&amount, "amount", "a", 0, "expense amount")
	addExpenseCmd.MarkFlagRequired("description")
	addExpenseCmd.MarkFlagRequired("category")
	addExpenseCmd.MarkFlagRequired("amount")
	RootCmd.AddCommand(addExpenseCmd)
}

func addExpense(cmd *cobra.Command, args []string) {
	controller.AddNewExpenseToList(description, category, amount)

}
