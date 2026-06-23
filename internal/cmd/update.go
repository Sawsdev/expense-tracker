package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	expenseIdToUpdate     int
	expenseNewDescription string
	expenseNewAmount      int
	updateExpenseCmd      = &cobra.Command{
		Use: "update",
		Short: "Update an existing expense",
		Long: "Update an existing expense, changing the --description and --amount by the given --id",
		Run: updateExpense,
	}
)

func init(){
	updateExpenseCmd.Flags().IntVarP(&expenseIdToUpdate, "id", "i", 0, "expense id to update")
	updateExpenseCmd.Flags().StringVarP(&expenseNewDescription, "description", "d", "", "new expense description")
	updateExpenseCmd.Flags().IntVarP(&expenseNewAmount, "amount", "a", 0, "new expense amount")
	RootCmd.AddCommand(updateExpenseCmd)
}

func updateExpense(cmd *cobra.Command, args [] string) {

	fmt.Println("Expense updated")
}
