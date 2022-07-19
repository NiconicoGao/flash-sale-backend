package utils

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Name        string `yaml:"name"`
	IP          string `yaml:"ip"`
	Port        string `yaml:"port"`
	UploadDir   string `yaml:"upload_dir"`
	ImagePrefix string `yaml:"image_prefix"`
}

type MysqlConfig struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Private  string `yaml:"private"`
	Db       string `yaml:"db"`
}

type RedisConfig struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

type KafkaConfig struct {
	IP        string `yaml:"ip"`
	Port      string `yaml:"port"`
	Topic     string `yaml:"topic"`
	GroupName string `yaml:"group_name"`
}

type ActiveMQConfig struct {
	Endpoint string `yaml:"endpoint"`
	Queue    string `yaml:"queue"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type MyConfig struct {
	ServerConfig   ServerConfig   `yaml:"server"`
	MysqlConfig    MysqlConfig    `yaml:"mysql"`
	RedisConfg     RedisConfig    `yaml:"redis"`
	KafkaConfig    KafkaConfig    `yaml:"kafka"`
	ActiveMQConfig ActiveMQConfig `yaml:"activemq"`
}

func WriteConfig(filename string, data *MyConfig) error {
	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, content, 0644)
	return err

}
func ReadConfig(filename string) (*MyConfig, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	myConfig := new(MyConfig)
	err = yaml.Unmarshal(content, myConfig)
	if err != nil {
		return nil, err
	}
	// decrypt, err := Decrypt(myConfig.MysqlConfig.Private, myConfig.MysqlConfig.Password)
	// if err != nil {
	// 	return nil, err
	// }

	// myConfig.MysqlConfig.Password = string(decrypt)

	if !strings.HasSuffix(myConfig.ServerConfig.ImagePrefix, "/") {
		myConfig.ServerConfig.ImagePrefix = myConfig.ServerConfig.ImagePrefix + "/"
	}

	return myConfig, nil
}
