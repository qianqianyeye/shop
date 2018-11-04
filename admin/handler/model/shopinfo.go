package model

type ShopInfo struct {
	ID int64 `gorm:"column:id" json:"id"`
	ShopName string `gorm:"column:shop_name" json:"shop_name"`
	RShopName string `gorm:"column:r_shop_name" json:"r_shop_name"`
	TypeId int64 `gorm:"column:type_id" json:"type_id"`
	ShopDescribe string `gorm:"column:shop_describe" json:"shop_describe"`
	RShopDescribe string `gorm:"column:r_shop_describe" json:"r_shop_describe"`
	MarketPrice float64 `gorm:"column:market_price" json:"market_price"`
	DiscountPrice float64 `gorm:"column:discount_price" json:"discount_price"`
	RMarketPrice float64 `gorm:"column:r_market_price" json:"r_market_price"`
	RDiscountPrice float64 `gorm:"column:r_discount_price" json:"r_discount_price"`
	ContactType int64 `gorm:"column:contact_type" json:"contact_type"`
	ContactInfo string `gorm:"column:contact_info" json:"contact_info"`
	SortWeight int64 `gorm:"column:sort_weight" json:"sort_weight"`
	CreateAt string `gorm:"column:create_at" json:"create_at"`
	UpdateAt string `gorm:"column:update_at" json:"update_at"`

	ShopStyle []ShopStyle `gorm:"ForeignKey:ShopId;AssociationForeignKey:ID;" json:"shop_style,omitempty"`
	Image []Image `gorm:"ForeignKey:TargetId;AssociationForeignKey:ID;" json:"image,omitempty"`
	ShopType []ShopType `gorm:"ForeignKey:ID;AssociationForeignKey:TypeId;" json:"shop_type,omitempty"`
}

func (ShopInfo) TableName() string {
	return "shop_info"
}