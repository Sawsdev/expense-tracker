package budgetlist

import (
	"log"
	"time"

	"github.com/sawsdev/expense-tracker/internal/budget"
)

const (
	dateLayout = "2006-01-02"
)

type BudgetList struct {
	Budgets []budget.Budget
}

func NewBudgetList() BudgetList {
	return BudgetList{
		Budgets: []budget.Budget{},
	}
}

func AddNewBudget(budgetList *BudgetList, id int, budgetType int, cap int, description string, dateStart string, dateEnd string) {
	formattedDateStart, err := time.Parse(dateLayout, dateStart)
	if err != nil {
		log.Fatal("Error parsing start date", err)
	}

	formattedDateEnd, err := time.Parse(dateLayout, dateEnd)
	if err != nil {
		log.Fatal("Error parsing end date", err)
	}

	newBudget := budget.NewBudget(id, budgetType, cap, description, formattedDateStart.Format(dateLayout), formattedDateEnd.Format(dateLayout))

	budgetList.Budgets = append(budgetList.Budgets, newBudget)

}
