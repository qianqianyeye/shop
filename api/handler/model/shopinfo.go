package model

type ShopInfo struct {
	ID int64 `gorm:"column:id" json:"id,omitempty"`
	ShopName string `gorm:"column:shop_name" json:"shop_name,omitempty"`
	RShopName string `gorm:"column:r_shop_name" json:"r_shop_name,omitempty"`
	TypeId int64 `gorm:"column:type_id" json:"type_id,omitempty"`
	ShopDescribe string `gorm:"column:shop_describe" json:"shop_describe,omitempty"`
	RShopDescribe string `gorm:"column:r_shop_describe" json:"r_shop_describe,omitempty"`
	MarketPrice float64 `gorm:"column:market_price" json:"market_price,omitempty"`
	DiscountPrice float64 `gorm:"column:discount_price" json:"discount_price,omitempty"`
	RMarketPrice float64 `gorm:"column:r_market_price" json:"r_market_price,omitempty"`
	RDiscountPrice float64 `gorm:"column:r_discount_price" json:"r_discount_price,omitempty"`
	ContactType int64 `gorm:"column:contact_type" json:"contact_type,omitempty"`
	ContactInfo string `gorm:"column:contact_info" json:"contact_info,omitempty"`
	SortWeight int64 `gorm:"column:sort_weight" json:"sort_weight,omitempty"`
	CreateAt string `gorm:"column:create_at" json:"create_at,omitempty"`
	UpdateAt string `gorm:"column:update_at" json:"update_at,omitempty"`

	ShopStyle []ShopStyle `gorm:"ForeignKey:ShopId;AssociationForeignKey:ID;" json:"shop_style,omitempty"`
	Image []Image `gorm:"ForeignKey:TargetId;AssociationForeignKey:ID;" json:"image,omitempty"`
	ShopType []ShopType `gorm:"ForeignKey:ID;AssociationForeignKey:TypeId;" json:"shop_type,omitempty"`
}

func (ShopInfo) TableName() string {
	return "shop_info"
}