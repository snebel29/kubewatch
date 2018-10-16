package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"testing"
)

var kubeWatchDir string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	kubeWatchDir = path.Join(path.Dir(filename), "fixtures")
}

func TestInitConfigWithExplicitFileWorks(t *testing.T) {
	err := InitConfig(&InitArgs{
		ConfigFile:     path.Join(kubeWatchDir, ".kubewatch.yaml"),
		ConfigDir:      "",
		ConfigFileName: "",
	})
	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
	viper.Reset()
}

func TestInitConfigWithExplicitFileFail(t *testing.T) {
	err := InitConfig(&InitArgs{
		ConfigFile:     path.Join(kubeWatchDir, ".non-existent.yaml"),
		ConfigDir:      "",
		ConfigFileName: "",
	})
	if err == nil {
		t.Errorf("Failed with error: %s", err)
	}
	viper.Reset()
}

func TestConfigWithConfigDiscoveryWorks(t *testing.T) {
	err := InitConfig(&InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".kubewatch",
	})
	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
	viper.Reset()
}

func TestConfigWithConfigDiscoveryFail(t *testing.T) {
	err := InitConfig(&InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".non-existent",
	})
	if err == nil {
		t.Errorf("Failed with error: %s", err)
	}
	viper.Reset()
}

func TestLoadConfigIsParsedCorrectlyWithConfigDir(t *testing.T) {
	InitConfig(&InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".kubewatch",
	})
	cfg, _ := NewConfig()
	if cfg.Resource.Deployment != true {
		t.Errorf("Failed with value %t", cfg.Resource.Deployment)
	}
	if cfg.Namespace != "default" {
		t.Errorf("Failed with value %s", cfg.Namespace)
	}
	viper.Reset()
}

func TestLoadConfigIsParsedCorrectlyWithConfigFile(t *testing.T) {
	InitConfig(&InitArgs{
		ConfigFile:     path.Join(kubeWatchDir, ".kubewatch.yaml"),
		ConfigDir:      "",
		ConfigFileName: "",
	})
	cfg, _ := NewConfig()
	if cfg.Resource.Deployment != true {
		t.Errorf("Failed with value %t", cfg.Resource.Deployment)
	}
	if cfg.Namespace != "default" {
		t.Errorf("Failed with value %s", cfg.Namespace)
	}
	viper.Reset()
}

func TestLoadConfigFailWithBadYaml(t *testing.T) {
	err := InitConfig(&InitArgs{
		ConfigFile:     path.Join(kubeWatchDir, ".empty-kubewatch.yaml"),
		ConfigDir:      "",
		ConfigFileName: "",
	})

	_, err = NewConfig()
	if err == nil {
		t.Error("Failed without error")
	}
	viper.Reset()
}

func TestConfigWithAutomaticEnvWorks(t *testing.T) {
	os.Setenv("KW_RESOURCE_DEPLOYMENT", "false")
	InitConfig(&InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".kubewatch",
	})

	cfg, _ := NewConfig()
	if cfg.Resource.Deployment != false {
		t.Errorf("Failed with value %t", cfg.Resource.Deployment)
	}
	viper.Reset()
}
