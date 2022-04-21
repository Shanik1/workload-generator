/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/shanik1/workload-generator/pkg/deployer"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete workloads created by workload-generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		deleteWorkloads()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteWorkloads() {
	workloadType := normalizeWorkloadType(workloadSettings.WorkloadType)
	workloadDeployer, err := deployer.NewWorkloadsDeployer(workloadType, workloadSettings.WorkloadName, workloadSettings.KubeConfigPath, workloadSettings.Namespace)
	if err != nil {
		logrus.Errorf("could not generate kubernetes client: %v", err)
		return
	}
	if err := workloadDeployer.DeleteWorkloads(); err != nil {
		logrus.Errorf("failed deleting workload-generator workloads: %v", err)
	} else {
		logrus.Infof("workloads deleted in namespace %s", workloadSettings.Namespace)
	}
}
