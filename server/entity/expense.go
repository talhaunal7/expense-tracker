package entity

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserID      int
	CategoryID  uint
	Amount      float64 `gorm:"not null"`
	Description string
	/*ExpenseDate time.Time `gorm:"not null"` */
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User     User
	Category Category
}
