package main


import (
	"os"
	"strings"
	"errors"
	"log"
	"github.com/BurntSushi/toml"
	
	_ "github.com/awslabs/aws-sdk-go/aws"
	_ "github.com/awslabs/aws-sdk-go/gen/ec2"	
)

const DefaultConfigFile = "aws_keys.toml"

type ConfToml struct {
	Keys SectionKeys `toml:"keys"`
}
type SectionKeys struct {
	Access_key string `toml:"access_key"`
	Secret_key string `toml:"secret_key"`
}

// user home dir
func getUserHomeDir() string {
	return os.Getenv("HOME")
}

// uesr home dir replace
func replaceHomeDir(confPath string) string {
	return strings.Replace(confPath, "~", getUserHomeDir(), 1)
}

// config の パース
func SetConf() {
	confPath := getUserHomeDir() + "/" + DefaultConfigFile
	err := LoadConf(confPath, &conftoml)
	if err != nil {
		log.Fatal(err)
	}

}

func LoadConf(path string, conftoml *ConfToml) error {
	_, err := toml.DecodeFile(path, &conftoml)
	if err != nil {
		panic(err)
	}

	// error チェック
	if conftoml.Keys.Access_key == "" {
		return errors.New("Consumer key is empty")
	}

	if conftoml.Keys.Secret_key == "" {
		return errors.New("Access token is empty")
	}

	return nil

}

func init() {
	SetConf()			
}
