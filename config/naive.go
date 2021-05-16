package config

import "fmt"

type naiveConfigManager struct {
	config map[string]interface{}
}

func (n *naiveConfigManager) Type() ConfigManagerType {
	return CONFIG_NAIVE
}

func (n *naiveConfigManager) SetValue(key string, value interface{}) error {
	n.config[key] = value
	return nil
}

func (n *naiveConfigManager) GetString(key string) string {
	if value, exists := n.config[key]; exists {
		return fmt.Sprintf("%s", value)
	} else {
		return ""
	}
}

func (n *naiveConfigManager) SetMap(m map[string]interface{}) error {
	for k,v := range m {
		n.SetValue(k, v)
	}
	return nil
}
