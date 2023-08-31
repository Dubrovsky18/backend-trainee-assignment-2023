package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id    int      `json:"id,omitempty" gorm:"column:id"`
	Slugs []string `json:"slug" gorm:"-"`
}

type Slug struct {
	gorm.Model
	NameSlug string `json:"name_slug" gorm:"primaryKey"`
}

type AddRemoveUserSlug struct {
	AddSegments    []string `json:"add_segments,omitempty"`
	RemoveSegments []string `json:"del_segments,omitempty"`
}

type UserSegmentHistory struct {
	gorm.Model
	UserId      int
	SegmentSlug string
	Operation   string
	Timestamp   time.Time
}

type RelationUserSlug struct {
	gorm.Model
	UserId   int    `gorm:"column:user_id"`
	NameSlug string `gorm:"column:name_slug"`
}

type UserHistory struct {
	UserId      int `json:"user_id,omitempty"`
	YearStart   int `json:"year_start"`
	YearFinish  int `json:"year_finish"`
	MonthStart  int `json:"month_start"`
	MonthFinish int `json:"month_finish"`
}
