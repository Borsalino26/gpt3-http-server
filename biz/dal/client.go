package dal

import (
	"fmt"
	"github.com/goccy/go-json"
	goGPT "github.com/sashabaranov/go-gpt3"
	"os"
)

var GPTClient *goGPT.Client

func init() {
	config, err := ReadServerConfig()
	if err != nil {
		panic(err)
	}
	GPTClient = goGPT.NewClient(config.GPTAuthToken)
}

func ReadServerConfig() (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile("config.json")
	if err != nil {
		return config, fmt.Errorf("read config.json error: %v", err)
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		return config, fmt.Errorf("read config.json error: %v", err)
	}
	return config, nil
}
