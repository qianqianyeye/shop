package model

type ShopStyle struct {
	ID int64 `gorm:"column:id" json:"id"`
	StyleName string `gorm:"column:style_name" json:"style_name"`
	StyleDescribe string `gorm:"column:style_describe" json:"style_describe"`
	ShopId int64 `gorm:"column:shop_id" json:"shop_id"`
	RStyleName string `gorm:"column:r_style_name" json:"r_style_name"`
	RStyleDescribe string `gorm:"column:r_style_describe" json:"r_style_describe"`
	Newest int  `gorm:"column:newest" json:"newest"`
	Image []Image `gorm:"ForeignKey:TargetId;AssociationForeignKey:ID;" json:"image"`
}

func (ShopStyle) TableName() string {
	return "shop_style"
}