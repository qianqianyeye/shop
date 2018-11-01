package model

type Admin struct {
	ID int64 `gorm:"column:id" json:"id"`
	RoleId int64 `gorm:"column:role_id" json:"role_id"`
	Name string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
	Contacts string `gorm:"column:contacts" json:"contacts"`
	Phone string `gorm:"column:phone" json:"phone"`
	Status int64 `gorm:"column:status" json:"status"`
	Email string `gorm:"column:email" json:"email"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func (Admin) TableName() string {
	return "admin"
}