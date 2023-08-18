package models

import (
	"time"
	// "gorm.io/gorm"
)

type Video struct {
	// gorm.Model // 包含Id、CreatedAt、UpdatedAt
	Id            int64     `json:"id,omitempty" gorm:"column:id"`
	Title         string    `json:"title,omitempty" gorm:"column:title"`
	AuthorID      int64     `json:"author_id,omitempty" gorm:"column:author_id"`
	Author        User      `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	PlayUrl       string    `json:"play_url,omitempty" gorm:"column:paly_url"`
	CoverUrl      string    `json:"cover_url,omitempty" gorm:"column:cover_url"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	// UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func AddVideo(video *Video) error {
	return DB.Create(&video).Error
}
