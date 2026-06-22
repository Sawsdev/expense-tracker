package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra")

var RootCmd = &cobra.Command{
	Use: "expense-tracker",
	Short: "Expense tracker is an app for keep track of personal expenses",
	Long: `Simple expense tracker with date details and filters by month, built in golang for serge345 and roadmap.sh task.
	Complete docs are found in: https://github.com/Sawsdev/expense-tracker 
	`,

}

func Execute(){
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init(){
}

func ShowWelcomeMessage(){
	fmt.Print(`** Use the command "add --description 'yourdescriptionhere' --amount 30" to add new expense with a given amount and description.
	** Use the command "list" to view full expense list
	** Use the command "summary" to show a total expense summary
	** Use the command "summary --month 9" to show a total expense summary by a given month
	** Use the command "delete --id 1" to delete a given expense by expense id
	** Use the command "update --id 1 --description 'new description' --amount 30" to update a given expense by expense id`)
}