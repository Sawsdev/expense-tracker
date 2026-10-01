package cmd

import (
	"github.com/sawsdev/expense-tracker/internal/controller"
	"github.com/spf13/cobra"
)

var (
	expenseIdToUpdate     int
	expenseNewDescription string
	expenseNewCategory    string
	expenseNewAmount      int
	expenseNewDate        string
	updateExpenseCmd      = &cobra.Command{
		Use:   "update",
		Short: "Update an existing expense",
		Long:  "Update an existing expense, changing the --description and --amount by the given --id",
		Run:   updateExpense,
	}
)

func init() {
	updateExpenseCmd.Flags().IntVarP(&expenseIdToUpdate, "id", "i", 0, "expense id to update")
	updateExpenseCmd.Flags().StringVarP(&expenseNewDescription, "description", "d", "", "new expense description")
	updateExpenseCmd.Flags().StringVarP(&expenseNewCategory, "category", "c", "", "new expense category")
	updateExpenseCmd.Flags().StringVarP(&expenseNewDate, "date", "D", "", "new expense date")
	updateExpenseCmd.Flags().IntVarP(&expenseNewAmount, "amount", "a", 0, "new expense amount")
	updateExpenseCmd.MarkFlagRequired("id")
	updateExpenseCmd.MarkFlagRequired("description")
	updateExpenseCmd.MarkFlagRequired("category")
	updateExpenseCmd.MarkFlagRequired("date")
	updateExpenseCmd.MarkFlagRequired("amount")
	RootCmd.AddCommand(updateExpenseCmd)
}

func updateExpense(cmd *cobra.Command, args []string) {

	controller.UpdateExpense(expenseIdToUpdate, expenseNewDescription, expenseNewCategory, expenseNewDate, expenseNewAmount)
}
