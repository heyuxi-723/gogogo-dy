package models

type User struct {
	Id              int64  `json:"id,omitempty" gorm:"column:id"`
	Name            string `json:"name,omitempty" gorm:"column:name"`
	Password        string `json:"password,omitempty" gorm:"column:password"`
	Avatar          string `json:"avatar,omitempty" gorm:"column:avatar"`
	BackgroundImage string `json:"background_image,omitempty" gorm:"column:background_image"`
	Signature       string `json:"signature,omitempty" gorm:"column:signature"`
}

func AddUser(user *User) error {
	return DB.Create(&user).Error
}
func IsUserExistByUsername(username string) bool {
	var user User
	DB.Where("name = ?", username).First(&user)
	if user.Id == 0 {
		return false
	}
	return true
}
