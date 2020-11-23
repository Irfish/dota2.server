package base

import (
	"encoding/json"
	"github.com/Irfish/component/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    		string
	LogPath     		string
	LogFlag     		string
	EtcdAddr    		string
	GinAddr     		string
	PayPayClientID   	string
	PayPaySecret     	string
	PayPayUrl     	    string
	PayPayVersion     	string
}

type LotteryConfig struct{
	Id int
	Type int
	Value int
	Rate  int
}

var SilverLotteryConfig []LotteryConfig
var GoldLotteryConfig []LotteryConfig

func init() {
	data, err := ioutil.ReadFile("config/login.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}

	data, err = ioutil.ReadFile("config/silver_lottery.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &SilverLotteryConfig)
	if err != nil {
		log.Fatal("%v", err)
	}

	for k,v:=range SilverLotteryConfig {
	  log.Debug("SilverLotteryConfig: %d   %v",k,v)
	}

	data, err = ioutil.ReadFile("config/gold_lottery.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &GoldLotteryConfig)
	if err != nil {
		log.Fatal("%v", err)
	}

	for k,v:=range GoldLotteryConfig {
		log.Debug("GoldLotteryConfig: %d   %v",k,v)
	}

}
