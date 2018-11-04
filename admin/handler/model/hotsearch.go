package model

type HotSearch struct {
	ID int64 `gorm:"column:id" json:"id"`
	HotName string `gorm:"column:hot_name" json:"hot_name"`
	RHotName string `gorm:"column:r_hot_name" json:"r_hot_name"`
	CreateAt string `gorm:"column:create_at" json:"create_at"`
	UpdateAt string `gorm:"column:update_at" json:"update_at"`
	Sort int64 `gorm:"column:sort" json:"sort"`

	Action int  `gorm:"column:action" json:"action"`//1上移，2下移
}
func (HotSearch) TableName() string {
	return "hot_search"
}