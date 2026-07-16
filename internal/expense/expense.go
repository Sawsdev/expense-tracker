package expense

/**
* Expense fields
* Id - int
* Date - DateFormat: YYYY-MM-DD
* Description - string
* Amount - int(usd)
*
*/

type Expense struct {
	Id int `json:"id"`
	Date string `json:"date"`
	Description string `json:"description"`
	Amount int `json:"amount"`
}

func NewExpense(id int, description string, date string, amount int) Expense {
	return Expense{
		Id: id,
		Date: date,
		Description: description,
		Amount: amount,
	}
}
