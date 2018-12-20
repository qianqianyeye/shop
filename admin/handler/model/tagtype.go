package model

type TagType struct {
	ID int64 `gorm:"column:id" json:"id"`
	ShopTypeId int64 `gorm:"column:shop_type_id" json:"shop_type_id"`
	TagId int64 `gorm:"column:tag_id" json:"tag_id"`
	ShopInfoId int64 `gorm:"column:shop_info_id" json:"shop_info_id"`
}

func (TagType) TableName() string {
	return "tag_type"
}