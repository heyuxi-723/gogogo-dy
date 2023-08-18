package models

import (
	"time"
)

type Comment struct {
	// gorm.Model
	Id      int64  `json:"id,omitempty"`
	UserID  int64  `json:"-" gorm:"column:user_id"`
	User    User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	VideoID int64  `json:"-" gorm:"column:video_id"`
	Video   Video  `json:"-" gorm:"foreignKey:VideoID"`
	Content string `json:"content,omitempty" gorm:"column:content"`
	// CreateDate string `json:"create_date,omitempty"`
	CreatedAt time.Time `json:"created_date,omitempty"`
}

func AddComment(comment *Comment) error {
	return DB.Create(comment).Error
}

func DelComment(commentID int64) error {
	return DB.Delete(&Comment{}, commentID).Error
}
