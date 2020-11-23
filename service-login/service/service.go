package service

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/component/xorm"
	"github.com/Irfish/dota2.server/service-login/base"
	mGin "github.com/Irfish/dota2.server/service-login/gin"
	log1 "log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	logger, err := log.New(base.Server.LogLevel, base.Server.LogPath, log1.LstdFlags)
	if err != nil {
		panic(err)
	}
	log.Export(logger)
	defer logger.Close()

	log.Debug("fantasy service login running ")
	redis.Run()
	xorm.Run()
	mGin.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}
