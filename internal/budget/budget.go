package budget

/**
* Budget fields
* Id - int
* Type - Enum string ["day", "week", "month"]
* Cap - int
* Description - string
* DateStart - DateFormat: YYYY-MM-DD
* DateEnd - DateFormat: YYYY-MM-DD
 */

type BudgetType int

const (
	Daily = iota
	Weekly
	Monthly
)

var budgetTypeName = map[BudgetType]string{
	Daily:   "daily",
	Weekly:  "weekly",
	Monthly: "monthly",
}

type Budget struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Cap         int    `json:"cap"`
	Description string `json:"description"`
	DateStart   string `json:"dateStart"`
	DateEnd     string `json:"dateEnd"`
}

func NewBudget(id int, budgetType int, cap int, description string, dateStart string, dateEnd string) Budget {
	return Budget{
		Id:          id,
		Type:        budgetTypeName[BudgetType(budgetType)],
		Cap:         cap,
		Description: description,
		DateStart:   dateStart,
		DateEnd:     dateEnd,
	}
}
