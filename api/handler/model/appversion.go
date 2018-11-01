package model

type AppVersion struct {
	ID int64 `gorm:"column:id" json:"id"`
	VersionNum string `gorm:"column:version_num" json:"version_num"`
	VersionDescript string `gorm:"column:version_descript" json:"version_descript"`
	Type string `gorm:"column:type" json:"type"`
	DownAddr string `gorm:"column:down_addr" json:"down_addr"`
	CreateAt string `gorm:"column:create_at" json:"create_at"`
	UpdateAt string `gorm:"column:update_at" json:"update_at"`
}
func (AppVersion) TableName() string {
	return "app_version"
}