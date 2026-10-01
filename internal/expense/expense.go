package expense

/**
* Expense fields
* Id - int
* Date - DateFormat: YYYY-MM-DD
* Description - string
* Caregory - string
* Amount - int(usd)
*
 */

type Expense struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
}

func NewExpense(id int, date string, description string, category string, amount int) Expense {
	return Expense{
		Id:          id,
		Date:        date,
		Description: description,
		Category:    category,
		Amount:      amount,
	}
}
