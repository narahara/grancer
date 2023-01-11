package model

import "time"

type PriceHistory struct {
	TradeCode string    `gorm:"size:16;not null;primary_key"`
	TimeFrame string    `gorm:"size:2;not null;primary_key"`
	DateTime  time.Time `gorm:"not null;primary_key"`
	Open      float32   `gorm:"not null;default:0.0"`
	High      float32   `gorm:"not null;default:0.0"`
	Low       float32   `gorm:"not null;default:0.0"`
	Close     float32   `gorm:"not null;default:0.0"`
	Volume    float32   `gorm:"not null;default:0.0"`
	UpdatedAt time.Time `gorm:"not null"`
}
