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
	"errors"
	"fmt"
	"github.com/snebel29/kubewatch/pkg/logging"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Handler   Handler  `mapstructure:"handler"`
	Resource  Resource `mapstructure:"resource"`
	Namespace string   `mapstructure:"namespace,omitempty"`
	Log       *logging.Logger
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
	ReplicationController bool `mapstructure:"replicationcontroller"`
	ReplicaSet            bool `mapstructure:"replicaset"`
	DaemonSet             bool `mapstructure:"daemonset"`
	Services              bool `mapstructure:"services"`
	Pod                   bool `mapstructure:"pod"`
	Job                   bool `mapstructure:"job"`
	PersistentVolume      bool `mapstructure:"persistentvolume"`
	Namespace             bool `mapstructure:"namespace"`
	Secret                bool `mapstructure:"secret"`
	ConfigMap             bool `mapstructure:"configmap"`
	Ingress               bool `mapstructure:"ingress"`
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
	Channel  string `mapstructure:"channel"`
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
}

type Flock struct {
	Url string `mapstructure:"url"`
}

type Webhook struct {
	Url string `mapstructure:"url"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	viper.Unmarshal(cfg)
	if reflect.DeepEqual(cfg, &Config{}) {
		return cfg, errors.New("Unmarshaled config equals &Config")
	} else {
		return cfg, nil
	}
}

type InitArgs struct {
	ConfigFile     string
	ConfigDir      string
	ConfigFileName string // Without extension!!
}

// TODO: Do we really need to initialize separately?
func InitConfig(args *InitArgs) error {
	replacer := strings.NewReplacer(".", "_")

	viper.SetEnvPrefix("KW")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	if args.ConfigFile != "" {
		if _, err := os.Stat(args.ConfigFile); err != nil {
			return errors.New(fmt.Sprintf("Failed to read config file %s", args.ConfigFile))
		}
		viper.SetConfigFile(args.ConfigFile)

	} else {
		viper.AddConfigPath(args.ConfigDir)
		viper.SetConfigName(args.ConfigFileName)
	}
	if err := viper.ReadInConfig(); err != nil {
		return errors.New(fmt.Sprintf("Reading config: %s", err))
	}
	return nil

}
