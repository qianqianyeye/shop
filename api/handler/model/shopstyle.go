package model

type ShopStyle struct {
	ID int64 `gorm:"column:id" json:"id"`
	StyleName string `gorm:"column:style_name" json:"style_name" binding:"required"`
	StyleDescribe string `gorm:"column:style_describe" json:"style_describe"`
	ShopId int64 `gorm:"column:shop_id" json:"shop_id"`
}

func (ShopStyle) TableName() string {
	return "shop_style"
}