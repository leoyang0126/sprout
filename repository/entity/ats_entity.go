package entity

import "time"

type AtsResumeBase struct {
	Id              int       `gorm:"column:id"`
	NameSpell       string    `gorm:"column:name_spell"`
	Phone           string    `gorm:"column:phone"`
	Email           string    `gorm:"column:email"`
	Gender          string    `gorm:"column:gender"`
	Age             string    `gorm:"column:age"`
	CurrentLocation string    `gorm:"column:current_location"`
	Ctime           time.Time `gorm:"column:ctime"`
	Mtime           time.Time `gorm:"column:mtime"`
	//Username     string  `gorm:"type:varchar(20);not null;unique;index:idx_username" json:"username"`
	//Password     string  `gorm:"size:255;not null" json:"password,omitempty"`
	//Mobile       string  `gorm:"type:varchar(11);not null;unique" json:"mobile"`
	//Avatar       string  `gorm:"type:varchar(255)" json:"avatar"`
	//Nickname     *string `gorm:"type:varchar(20)" json:"nickname"`
	//Introduction *string `gorm:"type:varchar(255)" json:"introduction"`
	//Status       uint    `gorm:"type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"`
	//Creator      string  `gorm:"type:varchar(20);" json:"creator"`
}
