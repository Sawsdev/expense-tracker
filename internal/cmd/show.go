package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
)

var(
	expenseId int
	day int
	month int
	year int
	showExpenseCmd = &cobra.Command{
		Use: "show",
		Short: "Show list of expenses use --id or --day & --month to filter",
		Long: "Show a full list of expenses, you can filter by an specific expense using --id or filter for a given date using mandatory flags --day --month --year ",
		Run: showExpense,
	}
)

func init() {
	showExpenseCmd.Flags().IntVarP(&expenseId, "id", "i", 0, "expense id to search")
	showExpenseCmd.Flags().IntVarP(&day, "day", "d", 1, "day of the expense to search")
	showExpenseCmd.Flags().IntVarP(&month, "month", "m", 1, "month of the expense to search for date")
	showExpenseCmd.Flags().IntVarP(&year, "year", "y", 1969, "year of the expense to search")
	RootCmd.AddCommand(showExpenseCmd)
}

func showExpense(cmd * cobra.Command, args [] string){
	fmt.Printf(`Expense found with id %d \n`, expenseId)
	fmt.Printf(`Expenses found in date: %d-%d-%d \n`, year,month,day)
}
