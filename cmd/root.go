/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dimchansky/utfbom"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "infinitas-streaming-tool",
	Short: "Useful stuff for Beatmania IIDX Infinitas streaming.",
	Long:  `Infinitas streaming stuff.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Infinitas Streaming tools starting...")

		var lastTitle = ""

		for range time.Tick(2 * time.Second) {

			// get title.txt and load into lastTitle variable
			file, err := os.Open("title.txt")
			if err == nil {
				scanner := bufio.NewScanner(file)
				scanner.Scan()
				titleOutput, titleError := ioutil.ReadAll(utfbom.SkipOnly(bytes.NewReader(scanner.Bytes())))
				if titleError == nil {
					lastTitle = string(titleOutput)
				}
			}
			defer file.Close()

			fmt.Println(lastTitle)
		}

	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.infinitas-go-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
