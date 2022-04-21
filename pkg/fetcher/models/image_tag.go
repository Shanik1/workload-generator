package models

import "time"

type ImageTagArchitectureDetails struct {
	Architecture string      `json:"architecture"`
	Features     string      `json:"features"`
	Variant      interface{} `json:"variant"`
	Digest       string      `json:"digest"`
	OS           string      `json:"os"`
	OSFeatures   string      `json:"os_features"`
	OSVersion    interface{} `json:"os_version"`
	Size         int64       `json:"size"`
	Status       string      `json:"status"`
	LastPulled   *time.Time  `json:"last_pulled"`
	LastPushed   *time.Time  `json:"last_pushed"`
}
type ImageTag struct {
	Creator       int                           `json:"creator"`
	ID            int                           `json:"id"`
	ImageID       string                        `json:"image_id"`
	Images        []ImageTagArchitectureDetails `json:"images"`
	LastUpdated   *time.Time                    `json:"last_updated"`
	Name          string                        `json:"name"`
	FullSize      int64                         `json:"full_size"`
	V2            bool                          `json:"v2"`
	TagStatus     string                        `json:"tag_status"`
	TagLastPulled *time.Time                    `json:"tag_last_pulled"`
	TagLastPushed *time.Time                    `json:"tag_last_pushed"`
}

type ImageTagList struct {
	Count   int        `json:"count"`
	Next    string     `json:"next"`
	Results []ImageTag `json:"results"`
}
