/*
Copyright 2016 Skippbox, Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/snebel29/kubewatch/config"
	kubewatch "github.com/snebel29/kubewatch/pkg/client"
	"github.com/spf13/cobra"
)

var configFileName, configFile string

var RootCmd = &cobra.Command{
	Use:   "kubewatch",
	Short: "Watch k8s events and trigger Handlers",
	Long:  "Watch k8s events and trigger Handlers",

	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		logrus.Debugf("Running with config: %#v", cfg)
		kubewatch.Run(cfg)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Fatalf("%s", err)
	}
}

func initConfig() {
	config.InitConfig(configFileName, configFile)
}

func init() {
	configFileName = ".kubewatch"
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(
		&configFile, "config", "", fmt.Sprintf("config file (default is ./%s)", configFileName),
	)
}
