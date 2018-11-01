package model

type ShopInfo struct {
	ID int64 `gorm:"column:id" json:"id"`
	ShopName string `gorm:"column:shop_name" json:"shop_name" binding:"required"`
	TypeId int64 `gorm:"column:type_id" json:"type_id" binding:"required"`
	ShopDescribe string `gorm:"column:shop_describe" json:"shop_describe"`
	MarketPrice float64 `gorm:"column:market_price" json:"market_price"`
	DiscountPrice float64 `gorm:"column:discount_price" json:"discount_price"`
	ContactType int64 `gorm:"column:contact_type" json:"contact_type" binding:"required"`
	ContactInfo string `gorm:"column:contact_info" json:"contact_info" binding:"required"`
	SortWeight int64 `gorm:"column:sort_weight" json:"sort_weight"`
	CreateAt string `gorm:"column:create_at" json:"create_at"`
	UpdateAt string `gorm:"column:update_at" json:"update_at"`
}

func (ShopInfo) TableName() string {
	return "shop_info"
}