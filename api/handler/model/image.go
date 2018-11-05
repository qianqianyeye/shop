package model

type Image struct {
	ID int64 `gorm:"column:id" json:"id,omitempty"`
	ImgUrl string `gorm:"column:img_url" json:"img_url,omitempty" binding:"required"`
	TargetId int64 `gorm:"column:target_id" json:"target_id,omitempty"`
	ImgType int64 `gorm:"column:img_type" json:"img_type,omitempty" binding:"required"`
	ImgName string `gorm:"column:img_name" json:"img_name,omitempty"`
	CreateAt string `gorm:"column:create_at" json:"create_at,omitempty"`
}

func (Image) TableName() string {
	return "image"
}