package gin

import (
	"encoding/json"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
)

func RedisParserGameRankManagerInit() {
	redis.AppendRedisMarshal(func(keys []string, i1 interface{}) (i interface{}, b bool) {
		if keys[0] == "hset" && keys[1] == RANK_LIST_REDIS_KEY && keys[2] == "info" {
			i0, ok := i1.(*GameRankManager)
			if ok {
				i2, e := json.Marshal(i0)
				if e != nil {
					log.Debug("RedisParserInit AppendRedisMarshal GameRankManager error:", e.Error())
					return
				}
				i = i2
				b = true
			}
		}
		return
	})
	redis.AppendRedisUnmarshal(func(keys []string, i1 interface{}) (i interface{}, b bool) {
		if keys[0] == "hget" && keys[1] == RANK_LIST_REDIS_KEY && keys[2] == "info" {
			i0, ok := i1.([]byte)
			if ok {
				gameRank := &GameRankManager{}
				e := json.Unmarshal(i0, gameRank)
				if e != nil {
					log.Debug("RedisParserInit AppendRedisUnmarshal GameRankManager error:", e.Error())
					return
				}
				i = gameRank
				b = true
			}
		}
		return
	})
}