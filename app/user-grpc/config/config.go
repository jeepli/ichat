package config

import (
	"flag"
	"log"

	commoncfg "github.com/jeepli/ichat/config"
)

var (
	gConf Config

	dbPathName      = "db.json"
	servicePathName = "service.json"

	configPath = flag.String("config", "", "configuration directory.")
)

type Config struct {
	Service commoncfg.ServiceConfig
	Db      commoncfg.DbConfig
}

func Conf() *Config {
	return &gConf
}

func Init() error {
	flag.Parse()

	err := commoncfg.ParseConfig(*configPath+"/"+servicePathName, &gConf.Service)
	if err != nil {
		log.Fatal("ParseConfig err ", err)
		return err
	}

	err = commoncfg.ParseConfig(*configPath+"/"+dbPathName, &gConf.Db)
	if err != nil {
		log.Fatal("ParseConfig err ", err)
		return err
	}

	return nil
}
