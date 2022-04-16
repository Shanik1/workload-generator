package models

import "time"

type ImageTagDetails struct {
}

type ImageTag struct {
	Creator       int            `json:"creator"`
	ID            int               `json:"id"`
	ImageID       string            `json:"image_id"`
	Images        []ImageTagDetails `json:"images"`
	LastUpdated   *time.Time        `json:"last_updated"`
	Name          string            `json:"name"`
	FullSize      int64             `json:"full_size"`
	V2            bool              `json:"v2"`
	TagStatus     string            `json:"tag_status"`
	TagLastPulled *time.Time        `json:"tag_last_pulled"`
	TagLastPushed *time.Time        `json:"tag_last_pushed"`
}

type ImageTagList struct {
	Count   int        `json:"count"`
	Next    string     `json:"next"`
	Results []ImageTag `json:"results"`
}
