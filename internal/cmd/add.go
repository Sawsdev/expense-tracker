package cmd

/**
* command form
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
**/

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addExpenseCmd)
}

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new expense entry",
	Long:  "Add a new expense entry to the list including --description and --amount",
	Run:   addExpense,
}

func addExpense(cmd *cobra.Command, args []string) {
	fmt.Println("Expense added")
}
