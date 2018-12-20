package model

type Tag struct {
	ID int64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	RName string `gorm:"column:r_name" json:"r_name"`
	Time string `gorm:"column:time" json:"time"`
	//ShopType []ShopType   `gorm:"many2many:tag_type;" json:"shop_type,omitempty"`
	ShopInfo []ShopInfo   `gorm:"many2many:tag_type;" json:"shop_info,omitempty"`
}

func (Tag) TableName() string {
	return "tag"
}