package util

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	content, err := ioutil.ReadFile("../example.toml")
	if err != nil {
		t.Fatal(err)
	}
	var config ConfigMap
	err = json.Unmarshal(content, &config)
	if err != nil {
		t.Fatal(err)
	}

	var expectedConfig ConfigMap = ConfigMap{
		Host: "127.0.0.1",
		Port: 8000,
		Discord: DiscordConfigMap{
			Token: "your discord token",
			Guild: "discord guild",
		},
	}

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("config not loaded correctly: expected %+v, got %+v", expectedConfig, config)
	}
}
