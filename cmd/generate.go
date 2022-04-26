/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/shanik1/workload-generator/pkg/deployer"
	"github.com/shanik1/workload-generator/pkg/fetcher"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/rand"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate workloads",
	Long:  `Command to generate new workloads to your kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		generateWorkloads()
	},
}

var generateSettings struct {
	ReposCount int
	AllTags    bool
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().IntVarP(&generateSettings.ReposCount, "count", "c", 10, "amount of image repos to deploy")
	generateCmd.PersistentFlags().BoolVarP(&generateSettings.AllTags, "all-tags", "a", false, "should deploy all repo tags")
}

func generateWorkloads() {
	workloadType := normalizeWorkloadType(workloadSettings.WorkloadType)
	workloadDeployer, err := deployer.NewWorkloadsDeployer(workloadType, workloadSettings.WorkloadName, workloadSettings.KubeConfigPath, workloadSettings.Namespace)
	if err != nil {
		logrus.Errorf("could not generate kubernetes client: %v", err)
		return
	}

	imageFetcher := fetcher.NewImageFetcher(generateSettings.ReposCount)
	images, err := imageFetcher.FetchRandomImages()
	if err != nil {
		fmt.Println("error fetching image")
		return
	}

	for _, image := range images {
		if generateSettings.AllTags {
			for _, tag := range image.ImageTags.Results {
				deployWorkload(image.RepositoryMetadata.Name, tag.Name, workloadDeployer)
			}
		} else {
			imageTag := image.ImageTags.Results[rand.Intn(len(image.ImageTags.Results))]
			deployWorkload(image.RepositoryMetadata.Name, imageTag.Name, workloadDeployer)
		}
	}
}

func deployWorkload(repo string, tag string, workloadDeployer *deployer.WorkloadsDeployer) {
	fullTag := fmt.Sprintf("docker.io/%s:%s", repo, tag)
	if err := workloadDeployer.DeployWorkload(fullTag); err != nil {
		logrus.Errorf("failed deploying workload %v: %v", fullTag, err)
	}
}
