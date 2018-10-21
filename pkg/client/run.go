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

package client

import (
	"errors"
	"fmt"
	"reflect"
	"github.com/snebel29/kubewatch/config"
	"github.com/snebel29/kubewatch/pkg/controller"
	"github.com/snebel29/kubewatch/pkg/handlers"
	"github.com/snebel29/kubewatch/pkg/handlers/flock"
	"github.com/snebel29/kubewatch/pkg/handlers/hipchat"
	"github.com/snebel29/kubewatch/pkg/handlers/mattermost"
	"github.com/snebel29/kubewatch/pkg/handlers/slack"
	"github.com/snebel29/kubewatch/pkg/handlers/webhook"
)

// Run runs the event loop processing with given handler
func Run(conf *config.Config) error {
	eventHandler, err := getHandler(conf)
	if err != nil {
		return err
	}
	// TODO: Handle controller errors
	controller.Start(conf, eventHandler)
	return nil
}

func getHandler(c *config.Config) (handlers.Handler, error) {
	var eventHandler handlers.Handler
	var configuredHandlers []handlers.Handler

	if !reflect.DeepEqual(&c.Handler.Slack, &config.Slack{}) {
		configuredHandlers = append(configuredHandlers, new(slack.Slack))
	}
	if !reflect.DeepEqual(&c.Handler.Hipchat, &config.Hipchat{}) {
		configuredHandlers = append(configuredHandlers, new(hipchat.Hipchat))
	}
	if !reflect.DeepEqual(&c.Handler.Mattermost, &config.Mattermost{}) {
		configuredHandlers = append(configuredHandlers, new(mattermost.Mattermost))
	}
	if !reflect.DeepEqual(&c.Handler.Flock, &config.Flock{}) {
		configuredHandlers = append(configuredHandlers, new(flock.Flock))
	}
	if !reflect.DeepEqual(&c.Handler.Slack, &config.Webhook{}) {
		configuredHandlers = append(configuredHandlers, new(webhook.Webhook))
	}

	if len(configuredHandlers) != 1 {
		return nil, errors.New(fmt.Sprintf("You have to configure exactly one handler, instead got %d", len(configuredHandlers)))
	}

	eventHandler = configuredHandlers[0]
	if err := eventHandler.Init(c); err != nil {
		return nil, err
	}
	return eventHandler, nil
}
