package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/shanik1/workload-generator/pkg/fetcher/models"
	"io/ioutil"
	"net/http"
)

const (
	listImagesURL           = "https://hub.docker.com/api/content/v1/products/search?architecture=arm,arm64&image_filter=official&operating_system=linux&page_size=25&q=&type=image"
	imageContentMetadataURL = "https://hub.docker.com/api/content/v1/products/images/%s"
	repositoryMetadataURL   = "https://hub.docker.com/v2/repositories/library/%s/"
	imageTagsURL            = "https://hub.docker.com/v2/repositories/library/%s/tags/"
	imageTagURL             = "https://hub.docker.com/v2/repositories/library/%s/tags/%s"
)

type ImageFetcher struct {
	repoCount int
}

func NewImageFetcher(repoCount int) *ImageFetcher {
	return &ImageFetcher{repoCount: repoCount}
}

func (imageFetcher *ImageFetcher) FetchRandomImages() ([]*models.ImageMetadata, error) {
	fetchedImages, err := imageFetcher.ListImages(listImagesURL)
	if err != nil {
		return nil, err
	}

	images := make([]*models.ImageMetadata, 0)
	for len(images) < imageFetcher.repoCount {
		for _, image := range fetchedImages.Summaries {
			imageMetadata, err := imageFetcher.generateImageMetadata(image)
			if err != nil {
				continue
			}
			if imageMetadata != nil {
				images = append(images, imageMetadata)
			}
		}
		if fetchedImages.Next != "" {
			fetchedImages, err = imageFetcher.ListImages(fetchedImages.Next)
			if err != nil {
				return images, err
			}
		}
	}

	return images[0:imageFetcher.repoCount], err
}

func (imageFetcher *ImageFetcher) getAndBindResponse(url string, result interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching images")
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error fetching images")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &result)
}

func (imageFetcher *ImageFetcher) generateImageMetadata(imageSummary models.ImageSummary) (*models.ImageMetadata, error) {
	imageMetadata := &models.ImageMetadata{}
	imageContentMetadata, err := imageFetcher.FetchImageContentMetadata(imageSummary)
	if err != nil {
		return nil, err
	}
	imageMetadata.ImageContentMetadata = imageContentMetadata

	repositoryMetadata, err := imageFetcher.FetchRepositoryMetadata(imageSummary)
	if err != nil {
		return nil, err
	}
	imageMetadata.RepositoryMetadata = repositoryMetadata

	imageTags, err := imageFetcher.FetchImageTags(imageSummary)
	if err != nil {
		return nil, err
	}
	imageMetadata.ImageTags = imageTags
	return imageMetadata, nil
}

// ListImages list images in docker.io and try to fetch images to create workloads with
func (imageFetcher *ImageFetcher) ListImages(url string) (models.ImageListResponse, error) {
	var images models.ImageListResponse
	err := imageFetcher.getAndBindResponse(url, &images)
	return images, err
}

func (imageFetcher *ImageFetcher) FetchImageContentMetadata(imageSummary models.ImageSummary) (models.ImageContentMetadata, error) {
	url := fmt.Sprintf(imageContentMetadataURL, imageSummary.Slug)
	var imageContentMetadata models.ImageContentMetadata
	err := imageFetcher.getAndBindResponse(url, &imageContentMetadata)
	return imageContentMetadata, err
}

func (imageFetcher *ImageFetcher) FetchRepositoryMetadata(imageSummary models.ImageSummary) (models.RepositoryMetadata, error) {
	url := fmt.Sprintf(repositoryMetadataURL, imageSummary.Slug)
	var repositoryMetadata models.RepositoryMetadata
	err := imageFetcher.getAndBindResponse(url, &repositoryMetadata)
	return repositoryMetadata, err
}

func (imageFetcher *ImageFetcher) FetchImageTags(imageSummary models.ImageSummary) (models.ImageTagList, error) {
	url := fmt.Sprintf(imageTagsURL, imageSummary.Slug)
	var imageTags models.ImageTagList
	err := imageFetcher.getAndBindResponse(url, &imageTags)
	return imageTags, err
}

func (imageFetcher *ImageFetcher) FetchImageTag(repo, tag string) (models.ImageTag, error) {
	url := fmt.Sprintf(imageTagURL, repo, tag)
	var imageTag models.ImageTag
	err := imageFetcher.getAndBindResponse(url, &imageTag)
	return imageTag, err
}
