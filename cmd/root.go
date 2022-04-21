/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultNamespace = "default"

var cfgFile string

var workloadSettings struct {
	WorkloadType   string
	WorkloadName   string
	KubeConfigPath string
	Namespace      string
}

var (
	defaultKubeConfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "workload-generator",
	Short: "Tool to randomly generate k8s workloads",
	Long:  `Tool to randomly generate k8s workloads`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.workload-generator.yaml)")
	rootCmd.PersistentFlags().StringVar(&workloadSettings.WorkloadType, "workload-type", "Pod", "workload type to deploy (Pod, Deployment)")
	rootCmd.PersistentFlags().StringVar(&workloadSettings.WorkloadName, "workload-name", "deployed-workload", "workload prefix name to deploy")
	rootCmd.PersistentFlags().StringVarP(&workloadSettings.KubeConfigPath, "kubeconfig", "k", defaultKubeConfigPath, "cluster kubeconfig to apply workloads on")
	rootCmd.PersistentFlags().StringVarP(&workloadSettings.Namespace, "namespace", "n", defaultNamespace, "cluster namespace to apply workloads in")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".workload-generator" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".workload-generator")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
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
