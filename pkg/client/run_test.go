package client

import (
	"fmt"
	c "github.com/snebel29/kubewatch/config"
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

func TestConfigWithMoreThanOneHandlerShouldFail(t *testing.T) {
	err := c.InitConfig(&c.InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".kubewatch-two-handler-set",
	})
	if err != nil {
		t.Errorf("%s", err)
	}
	cfg, err := c.NewConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	_, err = getHandler(cfg)
	if err == nil {
		t.Error("Should have failed with more than one handler")
	}
	viper.Reset()
}

func TestConfigWithOneHandlerShouldWork(t *testing.T) {
	err := c.InitConfig(&c.InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".kubewatch-one-handler-set",
	})
	if err != nil {
		t.Errorf("%s", err)
	}
	cfg, err := c.NewConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	_, err = getHandler(cfg)
	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
	viper.Reset()
}