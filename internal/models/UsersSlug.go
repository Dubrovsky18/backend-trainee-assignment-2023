package models

import "gorm.io/gorm"

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

type RelationUserSlug struct {
	gorm.Model
	UserId   int    `gorm:"column:user_id"`
	NameSlug string `gorm:"column:name_slug"`
}
