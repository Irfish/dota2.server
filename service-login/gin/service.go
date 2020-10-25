package gin

import (
	g "github.com/Irfish/component/gin"
	"github.com/Irfish/dota2.server/service-login/base"
)

type Gin struct {
	Address string
}
func (p *Gin)Addr() string {
	return p.Address
}

func Run() {
	server := new(Gin)
	server.Address=base.Server.GinAddr
	g.Run(server)
	RedisParserInit()
	RedisDataLoad()
}

func RedisParserInit()  {
	RedisParserGameRankManagerInit()
	RedisParserPlayerLotteryInit()
}

func RedisDataLoad()  {
	LoadGameRankManager()
}
