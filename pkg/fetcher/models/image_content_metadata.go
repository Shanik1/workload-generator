package models

import "time"

type Repository struct {
	Namespace string `json:"namespace"`
	RepoName  string `json:"reponame"`
}

type Tag struct {
	Latest bool   `json:"latest"`
	Value  string `json:"value"`
}

type Version struct {
	Architecture string `json:"architecture"`
	OS           string `json:"os"`
	Tags         []Tag  `json:"tags"`
}

type Plan struct {
	ID                  string            `json:"id"`
	Name                string            `json:"name"`
	IsOffline           bool              `json:"is_offline"`
	Eusa                string            `json:"eusa"`
	EusaType            string            `json:"eusa_type"`
	Instructions        string            `json:"instructions"`
	OperatingSystems    []MetadataField   `json:"operating_systems"`
	Architectures       []MetadataField   `json:"architectures"`
	DownloadAttribute   string            `json:"download_attribute"`
	CertificationStatus string            `json:"certification_status"`
	DefaultVersion      map[string]string `json:"default_version"`
	Rank                int               `json:"rank"`
	Repositories        []Repository      `json:"repositories"`
	Versions            []Version         `json:"versions"`
}

type ImageContentMetadata struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	IsOffline        bool            `json:"is_offline"`
	CreatedAt        *time.Time      `json:"created_at"`
	UpdatedAt        *time.Time      `json:"updated_at"`
	Links            []interface{}   `json:"links"`
	OperatingSystems []MetadataField `json:"operating_systems"`
	Categories       []MetadataField `json:"categories"`
	Architectures    []MetadataField `json:"architectures"`
	FullDescription  string          `json:"full_description"`
	ShortDescription string          `json:"short_description"`
	Slug             string          `json:"slug"`
	Source           string          `json:"source"`
	Type             string          `json:"type"`
	Plans            []Plan          `json:"plans"`
}
