package model

type ShopType struct {
	ID int64 `gorm:"column:id" json:"id"`
	ParentId int64 `gorm:"column:parent_id" json:"parent_id"`
	CName string `gorm:"column:c_name" json:"c_name"`
	RName string `gorm:"column:r_name" json:"r_name"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdateAt string `gorm:"column:update_at" json:"update_at"`
	
	//TagId []int `gorm:"-" json:"tag_id,omitempty"`
	//Tags []*Tag `gorm:"many2many:tag_type;"json:"tag_type,omitempty"`
	Image []Image `gorm:"ForeignKey:TargetId;AssociationForeignKey:ID;" json:"image,omitempty"`
}


func (ShopType) TableName() string {
	return "shop_type"
}
