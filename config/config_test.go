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
		t.Errorf("Should have failed with error, instead got: %s", err)
	}
	viper.Reset()
}

func TestLoadConfigIsParsedCorrectly(t *testing.T) {
	err := InitConfig(&InitArgs{
		ConfigFile:     "",
		ConfigDir:      kubeWatchDir,
		ConfigFileName: ".kubewatch-all-set",
	})

	if err != nil {
		t.Errorf("%s", err)
	}

	cfg, _ := NewConfig()

	tests := []struct {
		option   string
		has      interface{}
		shouldBe interface{}
	}{
		{"handler.slack.channel", cfg.Handler.Slack.Channel, "default"},
		{"handler.slack.token", cfg.Handler.Slack.Token, "default"},
		{"handler.hipchat.token", cfg.Handler.Hipchat.Token, "default"},
		{"handler.hipchat.room", cfg.Handler.Hipchat.Room, "default"},
		{"handler.hipchat.url", cfg.Handler.Hipchat.Url, "default"},
		{"handler.mattermost.channel", cfg.Handler.Mattermost.Channel, "default"},
		{"handler.mattermost.url", cfg.Handler.Mattermost.Url, "default"},
		{"handler.mattermost.username", cfg.Handler.Mattermost.Username, "default"},
		{"handler.flock.url", cfg.Handler.Flock.Url, "default"},
		{"handler.webhook.url", cfg.Handler.Webhook.Url, "default"},

		{"deployment", cfg.Resource.Deployment, true},
		{"replicationcontroller", cfg.Resource.ReplicationController, true},
		{"replicaset", cfg.Resource.ReplicaSet, true},
		{"daemonset", cfg.Resource.DaemonSet, true},
		{"services", cfg.Resource.Services, true},
		{"pod", cfg.Resource.Pod, true},
		{"job", cfg.Resource.Job, true},
		{"persistentvolume", cfg.Resource.PersistentVolume, true},
		{"namespace", cfg.Resource.Namespace, true},
		{"secret", cfg.Resource.Secret, true},
		{"configmap", cfg.Resource.ConfigMap, true},
		{"ingress", cfg.Resource.Ingress, true},

		{"Namespace", cfg.Namespace, "namespace"},
	}

	for _, test := range tests {
		if test.has != test.shouldBe {
			t.Errorf("Option [%v] should be set to [%v] but has [%v]",
				test.option, test.shouldBe, test.has)
		}
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
