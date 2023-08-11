package model

type User struct {
	Id              int64  `json:"id,omitempty" gorm:"column:id"`
	Name            string `json:"name,omitempty" gorm:"column:name"`
	Password        string `json:"password,omitempty" gorm:"column:password"`
	Avatar          string `json:"avatar,omitempty" gorm:"column:avatar"`
	BackgroundImage string `json:"background_image,omitempty" gorm:"column:background_image"`
	Signature       string `json:"signature,omitempty" gorm:"column:signature"`
}
