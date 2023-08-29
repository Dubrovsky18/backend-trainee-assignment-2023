package models

type User struct {
	Id    int      `json:"id,omitempty" bson:"id,omitempty" db:"id"`
	Slugs []string `json:"slug" db:"name_slugs"`
}

type Slug struct {
	NameSlug string `json:"name_slug,omitempty" bson:"name_slug,omitempty" db:"name_slug"`
}

type AddRemoveUserSlug struct {
	AddSegments    []string `json:"add_segments,omitempty"`
	RemoveSegments []string `json:"remove_segments,omitempty"`
}
