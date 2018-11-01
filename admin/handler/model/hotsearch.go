package model

type HotSearch struct {
	ID int64 `gorm:"column:id" json:"id"`
	HotName string `gorm:"column:hot_name" json:"hot_name"`
	CreateName string `gorm:"column:create_name" json:"create_name"`
	UpdateName string `gorm:"column:update_name" json:"update_name"`
}
func (HotSearch) TableName() string {
	return "hot_search"
}