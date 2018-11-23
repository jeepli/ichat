package config

import (
	"encoding/json"
	"io/ioutil"
)

type ServiceConfig struct {
	Addr string
}

type DbConfig struct {
	Address            string
	User               string
	Password           string
	Database           string
	PoolSize           int
	PoolTimeout        int
	IdleTimeout        int
	IdleCheckFrequency int
}

func ParseConfig(path string, obj interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}
