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
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sort"
)

var generateSettings struct {
	WorkloadType   string
	WorkloadsCount int
	WorkloadName   string
	KubeConfigPath string
	Namespace      string
}

const defaultNamespace = "default"

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate workloads",
	Long:  `Command to generate new workloads to your kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		generateWorkloads()
	},
}

var (
	defaultKubeConfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
)

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.PersistentFlags().StringVar(&generateSettings.WorkloadType, "workload-type", "Pod", "workload type to deploy (Pod, Deployment)")
	rootCmd.PersistentFlags().StringVar(&generateSettings.WorkloadName, "workload-name", "deployed-workload", "workload prefix name to deploy")
	rootCmd.PersistentFlags().StringVarP(&generateSettings.KubeConfigPath, "kubeconfig", "k", defaultKubeConfigPath, "cluster kubeconfig to apply workloads on")
	rootCmd.PersistentFlags().StringVarP(&generateSettings.Namespace, "namespace", "n", defaultNamespace, "cluster namespace to apply workloads in")
}

func generateWorkloads() {
	imageFetcher := fetcher.NewImageFetcher()
	images, err := imageFetcher.FetchRandomImages()
	if err != nil {
		fmt.Println("error fetching image")
		return
	}
	workloadType := normalizeWorkloadType(generateSettings.WorkloadType)
	workloadDeployer, err := deployer.NewWorkloadsDeployer(workloadType, generateSettings.WorkloadName, generateSettings.KubeConfigPath, generateSettings.Namespace)
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

func normalizeWorkloadType(workloadType string) string {
	supportedWorkloads := []string{"Deployment", "Pod"}
	sort.Strings(supportedWorkloads)
	if index := sort.SearchStrings(supportedWorkloads, workloadType); index > len(supportedWorkloads) || index < 0 {
		return "Pod"
	}
	return workloadType
}
