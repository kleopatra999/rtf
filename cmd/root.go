// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	caseDir    string
	resultDir  string
	labels     string
	envVars    []string
	extra      bool
	loggerAddr string
	verbose    int
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           os.Args[0],
	Short:         "Run or provide information about local regression tests",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&caseDir, "casedir", "c", "cases", "Directory containing cases")
	RootCmd.PersistentFlags().StringVarP(&resultDir, "resultdir", "r", "_results", "Directory to place results in")
	RootCmd.PersistentFlags().StringVarP(&labels, "labels", "l", "", "Labels to apply (comma separated)")
	RootCmd.PersistentFlags().StringSliceVarP(&envVars, "envVars", "e", []string{}, "Add a environment variable (multiple uses allowed)")
	RootCmd.PersistentFlags().BoolVarP(&extra, "extra", "x", false, "Add extra debug info to log files")
	RootCmd.PersistentFlags().StringVar(&loggerAddr, "logger", "", "<host>:<port> of a socket to log stdout to")
	RootCmd.PersistentFlags().CountVarP(&verbose, "verbose", "v", "Increase verbosity level")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
