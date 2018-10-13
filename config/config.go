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

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Handler   Handler  `mapstructure:"handler"`
	Resource  Resource `mapstructure:"resource"`
	Namespace string   `mapstructure:"namespace,omitempty"`
	// for watching specific namespace, leave it empty for watching all.
	// this config is ignored when watching namespaces
}

type Handler struct {
	Slack      Slack      `mapstructure:"slack"`
	Hipchat    Hipchat    `mapstructure:"hipchat"`
	Mattermost Mattermost `mapstructure:"mattermost"`
	Flock      Flock      `mapstructure:"flock"`
	Webhook    Webhook    `mapstructure:"webhook"`
}

// Resource contains resource configuration
type Resource struct {
	Deployment            bool `mapstructure:"deployment"`
	ReplicationController bool `mapstructure:"rc"`
	ReplicaSet            bool `mapstructure:"rs"`
	DaemonSet             bool `mapstructure:"ds"`
	Services              bool `mapstructure:"svc"`
	Pod                   bool `mapstructure:"po"`
	Job                   bool `mapstructure:"job"`
	PersistentVolume      bool `mapstructure:"pv"`
	Namespace             bool `mapstructure:"ns"`
	Secret                bool `mapstructure:"secret"`
	ConfigMap             bool `mapstructure:"configmap"`
	Ingress               bool `mapstructure:"ing"`
}

type Slack struct {
	Token   string `mapstructure:"token"`
	Channel string `mapstructure:"channel"`
}

type Hipchat struct {
	Token string `mapstructure:"token"`
	Room  string `mapstructure:"room"`
	Url   string `mapstructure:"url"`
}

type Mattermost struct {
	Channel  string `mapstructure:"room"`
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
}

type Flock struct {
	Url string `mapstructure:"url"`
}

type Webhook struct {
	Url string `mapstructure:"url"`
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		logrus.Fatalf("Cannot unmarshal config: %s", err)
	}
	return cfg
}

func InitConfig(configFileName string, configFile string) {
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
