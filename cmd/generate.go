/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/shanik1/workload_generator/pkg/fetcher"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateWorkloads() {
	imageFetcher := fetcher.NewImageFetcher()
	images, err := imageFetcher.FetchRandomImages()
	if err != nil {
		fmt.Println("error fetching image")
		return
	}
	fmt.Println(images)
}
