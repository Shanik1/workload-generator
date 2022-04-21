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

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateWorkloads() {
	imageFetcher := fetcher.NewImageFetcher()
	images, err := imageFetcher.FetchRandomImages()
	if err != nil {
		fmt.Println("error fetching image")
		return
	}
	workloadType := normalizeWorkloadType(workloadSettings.WorkloadType)
	workloadDeployer, err := deployer.NewWorkloadsDeployer(workloadType, workloadSettings.WorkloadName, workloadSettings.KubeConfigPath, workloadSettings.Namespace)
	if err != nil {
		logrus.Errorf("could not generate kubernetes client: %v", err)
		return
	}

	for _, image := range images {
		fullTag := fmt.Sprintf("docker.io/%s:%s", image.RepositoryMetadata.Name, image.ImageTags.Results[0].Name)
		if err := workloadDeployer.DeployWorkload(fullTag); err != nil {
			logrus.Errorf("failed deploying workload %v: %v", fullTag, err)
		}
	}
}
