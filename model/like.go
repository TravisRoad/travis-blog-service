package model

import "gorm.io/gorm"

type LikeOp struct {
	gorm.Model
	Operation string `gorm:"column:operation" json:"operation"`
	Field     string `gorm:"column:field" json:"field"`
	UserAgent string `gorm:"column:user_agent" json:"user_agent"`
	IPAddress uint32 `gorm:"column:ip_address" json:"ip_address"`
}

type Like struct {
	gorm.Model
	Field  string `gorm:"column:field" json:"field"`
	Number int    `gorm:"column:number;default:0" json:"number"`
}
