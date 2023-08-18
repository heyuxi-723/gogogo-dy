package models

type User struct {
	Id              int64  `json:"id" gorm:"column:id"`
	Name            string `json:"name" gorm:"column:name"`
	Password        string `json:"password,omitempty" gorm:"column:password"`
	Avatar          string `json:"avatar,omitempty" gorm:"column:avatar"`
	BackgroundImage string `json:"background_image,omitempty" gorm:"column:background_image"`
	Signature       string `json:"signature" gorm:"column:signature"`
}

type UserInfo struct {
	User
	FollowCount    int64  `json:"follow_count,omitempty"`
	FollowerCount  int64  `json:"follower_count,omitempty"`
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited string `json:"total_favorited,omitempty"`
	WorkCount      int64  `json:"work_count,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
}

type Follow struct {
	//关注者
	FollowUserId int64 `json:"follow_userId,omitempty" gorm:"column:follow_userId"`
	//被关注者
	FollowerUserId int64 `json:"follower_userId,omitempty" gorm:"column:follower_userId"`
}

func AddUser(user *User) error {
	return DB.Create(&user).Error
}
func IsUserExistByUsername(username string) bool {
	user, ok := QueryUserLogin(username, "name")
	if !ok || user.Id == 0 {
		return false
	}
	return true
}

func QueryUserLogin(username string, key string) (User, bool) {
	var user User
	res := DB.Table("users").Where(key+" = ?", username)
	if key == "id" {
		res = res.Select("id", "name", "avatar", "background_image", "signature")
	}
	res = res.First(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return User{}, false
	}
	return user, true
}

func QueryIsFollow(userId string, myUserId string) bool {
	var follow Follow
	res := DB.Table("follows").Where("follow_userId = ? and follower_userId = ?", myUserId, userId).First(&follow)
	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}
