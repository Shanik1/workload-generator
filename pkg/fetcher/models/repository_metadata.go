package models

import "time"

type Permissions struct {
	Admin bool `json:"admin"`
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

type RepositoryMetadata struct {
	Affiliation       string      `json:"affiliation"`
	CanEdit           bool        `json:"can_edit"`
	CollaboratorCount int         `json:"collaborator_count"`
	Description       string      `json:"description"`
	FullDescription   string      `json:"full_description"`
	HasStarred        bool        `json:"has_starred"`
	HubUser           string      `json:"hub_user"`
	IsAutomated       bool        `json:"is_automated"`
	IsMigrated        bool        `json:"is_migrated"`
	IsPrivate         bool        `json:"is_private"`
	LastUpdated       *time.Time  `json:"last_updated"`
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	Permissions       Permissions `json:"permissions"`
	PullCount         int         `json:"pull_count"`
	RepositoryType    string      `json:"repository_type"`
	StarCount         int         `json:"star_count"`
	Status            int         `json:"status"`
	User              string      `json:"user"`
}
