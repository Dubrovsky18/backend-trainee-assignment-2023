package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id    int      `json:"id,omitempty" gorm:"column:user_id"`
	Slugs []string `json:"slug" gorm:"-"`
}

type Slug struct {
	gorm.Model
	Id       int    `gorm:"column:slug_id"`
	NameSlug string `json:"name_slug" db:"name_slug" gorm:"column:name_slug"`
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

type AddRemoveUserSlugWithTTL struct {
	AddRemoveUserSlug
	TTLMap map[string]time.Duration `json:"ttl_map,omitempty"`
}

type SlugWithAutoAddition struct {
	Slug                   Slug
	AutoAdditionPercentage int `json:"auto_addition_percentage,omitempty"`
}
