package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
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
	expenseSummaryCmd.Flags().IntVarP(&monthSummary, "month", "m", 1, "get summary by the given month")
	expenseSummaryCmd.Flags().IntVarP(&yearSummary, "year", "y", 1969, "get expense summary by the given year")
	RootCmd.AddCommand(expenseSummaryCmd)
}

func getExpenseSummary(cmd *cobra.Command, args []string) {
	fmt.Println("Total expenses: 5656")
}
