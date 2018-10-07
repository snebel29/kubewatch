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

type Config struct {
	Handler   Handler  `mapstructure:"handler"`
	Resource  Resource `mapstructure:"resource"`
	Namespace string   `mapstructure:"namespace,omitempty"`
	// for watching specific namespace, leave it empty for watching all.
	// this config is ignored when watching namespaces
}

type Handler struct {
	Slack      Slack      `json:"slack"`
	Hipchat    Hipchat    `json:"hipchat"`
	Mattermost Mattermost `json:"mattermost"`
	Flock      Flock      `json:"flock"`
	Webhook    Webhook    `json:"webhook"`
}

// Resource contains resource configuration
type Resource struct {
	Deployment            bool `json:"deployment"`
	ReplicationController bool `json:"rc"`
	ReplicaSet            bool `json:"rs"`
	DaemonSet             bool `json:"ds"`
	Services              bool `json:"svc"`
	Pod                   bool `json:"po"`
	Job                   bool `json:"job"`
	PersistentVolume      bool `json:"pv"`
	Namespace             bool `json:"ns"`
	Secret                bool `json:"secret"`
	ConfigMap             bool `json:"configmap"`
	Ingress               bool `json:"ing"`
}

type Slack struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

type Hipchat struct {
	Token string `json:"token"`
	Room  string `json:"room"`
	Url   string `json:"url"`
}

type Mattermost struct {
	Channel  string `json:"room"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type Flock struct {
	Url string `json:"url"`
}

type Webhook struct {
	Url string `json:"url"`
}
