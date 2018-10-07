/*
Copyright 2018 Bitnami

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

package webhook

import (
	"fmt"
	"os"

	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/snebel29/kubewatch/config"
	kbEvent "github.com/snebel29/kubewatch/pkg/event"
)

var webhookErrMsg = `
%s

You need to set Webhook url
using "--url/-u" or using environment variables:

export KW_WEBHOOK_URL=webhook_url

Command line flags will override environment variables

`

// Webhook handler implements handler.Handler interface,
// Notify event to Webhook channel
type Webhook struct {
	Url string
}

type WebhookMessage struct {
	Text string `json:"text"`
}

// Init prepares Webhook configuration
func (m *Webhook) Init(c *config.Config) error {
	url := c.Handler.Webhook.Url

	if url == "" {
		url = os.Getenv("KW_WEBHOOK_URL")
	}

	m.Url = url

	return checkMissingWebhookVars(m)
}

func (m *Webhook) ObjectCreated(obj interface{}) {
	notifyWebhook(m, obj, "created")
}

func (m *Webhook) ObjectDeleted(obj interface{}) {
	notifyWebhook(m, obj, "deleted")
}

func (m *Webhook) ObjectUpdated(oldObj, newObj interface{}) {
	notifyWebhook(m, newObj, "updated")
}

func notifyWebhook(m *Webhook, obj interface{}, action string) {
	e := kbEvent.New(obj, action)

	webhookMessage := prepareWebhookMessage(e, m)

	err := postMessage(m.Url, webhookMessage)
	if err != nil {
		logrus.Errorf("%s\n", err)
		return
	}

	logrus.Printf("Message successfully sent to %s at %s ", m.Url, time.Now())
}

func checkMissingWebhookVars(s *Webhook) error {
	if s.Url == "" {
		return fmt.Errorf(webhookErrMsg, "Missing Webhook url")
	}

	return nil
}

func prepareWebhookMessage(e kbEvent.Event, m *Webhook) *WebhookMessage {
	return &WebhookMessage{
		e.Message(),
	}

}

func postMessage(url string, webhookMessage *WebhookMessage) error {
	message, err := json.Marshal(webhookMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
