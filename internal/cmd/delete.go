package cmd

import(
	"fmt"
	"github.com/sawsdev/expense-tracker/internal/controller"
	"github.com/spf13/cobra"
)

var (
	expenseIdToDelete 		int
	deleteExpenseCmd		= &cobra.Command{
		Use: "delete",
		Short: "Delete an existing expense",
		Long: "Delete an existing expense by the given --id",
		Run:deleteExpense,
	}
)

func init(){
	deleteExpenseCmd.Flags().IntVarP(&expenseIdToDelete, "id", "i", 0, "expense id to delete")
	deleteExpenseCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(deleteExpenseCmd)

}

func deleteExpense(cmd *cobra.Command, args [] string){
	fmt.Println("Expense deleted")
	controller.DeleteExpense(expenseIdToDelete)
}