package model

type Image struct {
	ID int64 `gorm:"column:id" json:"id"`
	ImgUrl string `gorm:"column:img_url" json:"img_url" binding:"required"`
	TargetId int64 `gorm:"column:target_id" json:"target_id"`
	ImgType int64 `gorm:"column:img_type" json:"img_type" binding:"required"`
	ImgName string `gorm:"column:img_name" json:"img_name"`
	CreateAt string `gorm:"column:create_at" json:"create_at"`
}

func (Image) TableName() string {
	return "image"
}