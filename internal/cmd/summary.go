package cmd

import (

	"github.com/spf13/cobra"
	"github.com/sawsdev/expense-tracker/internal/controller"
)

var (
	daySummary       int
	monthSummary      int
	yearSummary       int
	expenseSummaryCmd = &cobra.Command{
		Use:   "summary",
		Short: "Get total expense summary",
		Long:  "Get total expense summary, it can be obtained by month using the --month flag or by year using the --year flag",
		Run:   getExpenseSummary,
	}
)

func init() {
	expenseSummaryCmd.Flags().IntVarP(&daySummary, "day", "d", 0, "get summary by the given month")
	expenseSummaryCmd.Flags().IntVarP(&monthSummary, "month", "m", 0, "get summary by the given month")
	expenseSummaryCmd.Flags().IntVarP(&yearSummary, "year", "y", 0, "get expense summary by the given year")
	RootCmd.AddCommand(expenseSummaryCmd)
}

func getExpenseSummary(cmd *cobra.Command, args []string) {
	controller.GetExpenseSummary(daySummary,monthSummary,yearSummary)
}
