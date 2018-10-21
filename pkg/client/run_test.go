package client

import (
	"fmt"
	"github.com/snebel29/kubewatch/config"
	"github.com/snebel29/kubewatch/pkg/handlers/webhook"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"testing"
)

var kubeWatchDir string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	kubeWatchDir = path.Join(path.Dir(filename), "fixtures")
	fmt.Println(kubeWatchDir)
}

func TestSelectHandlerFromConfigShouldWorks(t *testing.T) {
	if err := config.InitConfig(&config.InitArgs{
		ConfigFile: path.Join(kubeWatchDir, ".kubewatch.yaml"),
	}); err != nil {
		t.Fatal(err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		t.Error(err)
	}
	h, err := selectHandlerFromConfig(cfg)
	if err != nil {
		t.Error(err)
	}
	if _, ok := h.(*webhook.Webhook); !ok {
		t.Errorf("handler of type: %T\n", h)
	}
	viper.Reset()
}
