package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

type Conf struct {
	Mongo Mongo `yaml:"mongo"`
	Kafka Kafka `yaml:"kafka"`
	GRpc  GRpc  `yaml:"grpc"`
}

type Mongo struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

type Kafka struct {
	Brokers []string `yaml:"brokers"`
}

type GRpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var conf Conf
var mu sync.RWMutex

func init() {
	var filename string
	env := os.Getenv("RUN_ENV")
	if env != "" {
		filename = "config/config_" + env + ".yaml"
	} else {
		filename = "config/config_dev.yaml"
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("load config file %s failed, err: %+v", filename, err)
	}

	mu.Lock()
	defer mu.Unlock()

	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalf("unmarshal config file failed, err: %+v", err)
	}

	log.Info("load config file finished!")
}

func Config() *Conf {
	mu.RLock()
	defer mu.RUnlock()

	return &conf
}
