package entity

type Books struct {
	ID     uint   `json:"ID" gorm:"primary_key"`
	Title  string `gorm:"type:varchar(100)"`
	Author string `gorm:"type:varchar(100)"`
	Desc   string `gorm:"type:varchar(100)"`
}
