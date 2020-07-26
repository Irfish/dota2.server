package service

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/component/xorm"
	mGin "github.com/Irfish/dota2.server/service-login/gin"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	log.Debug("fantasy service login running ")
	redis.Run()
	xorm.Run()
	mGin.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}
