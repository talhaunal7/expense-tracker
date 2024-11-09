package request

type ExpenseAdd struct {
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"categoryID" binding:"required"`
}
