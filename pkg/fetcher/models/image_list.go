package models

type ImageSummary struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	Slug             string          `json:"slug"`
	OperatingSystems []MetadataField `json:"operating_systems"`
	Architectures    []MetadataField `json:"architectures"`
}

type ImageListResponse struct {
	Next      string         `json:"next"`
	Count     int            `json:"count"`
	Page      int            `json:"page"`
	Previous  string         `json:"previous"`
	PageSize  int            `json:"page_size"`
	Summaries []ImageSummary `json:"summaries"`
}
