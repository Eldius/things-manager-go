package model

import "time"

type Thing struct {
	ID             int             `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string          `json:"name" gorm:"uniqueIndex;unique;size:50"`
	Description    *string         `json:"description" gorm:"size:255"`
	Available      float64         `json:"available"`
	StockMovements []StockMovement `json:"movs"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type StockMovement struct {
	ID         int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Quantity   float64 `json:"quantity"`
	UnityValue float64 `json:"unityValue"`
}

type Batch struct {
	ID             int
	Name           string
	ExpirationDate time.Time
}
