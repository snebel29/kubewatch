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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/snebel29/kubewatch/config"
	kubewatch "github.com/snebel29/kubewatch/pkg/client"
)

var configFile, configFileName string

var RootCmd = &cobra.Command{
	Use:   "kubewatch",
	Short: "Watch k8s events and trigger Handlers",
	Long:  "Watch k8s events and trigger Handlers",

	Run: func(cmd *cobra.Command, args []string) {
		cfg := &config.Config{}
		if err := viper.Unmarshal(cfg); err != nil {
			logrus.Fatalf("Cannot unmarshal config: %s", err)
		}
		logrus.Infof("%+v", cfg)
		kubewatch.Run(cfg)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	configFileName = ".kubewatch"
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(
		&configFile, "config", "", fmt.Sprintf("config file (default is ./%s)", configFileName),
	)
}

func initConfig() {
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	replacer := strings.NewReplacer(".", "_")

	viper.SetEnvPrefix("KW")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.AddConfigPath(currentDir)
	viper.SetConfigName(configFileName)

	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Reading config: %s")
	}
}
