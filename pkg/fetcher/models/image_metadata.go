package models

type ImageMetadata struct {
	RepositoryMetadata   RepositoryMetadata
	ImageContentMetadata ImageContentMetadata
	ImageTags            ImageTagList
}
