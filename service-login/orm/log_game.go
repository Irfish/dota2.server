package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type LogGame struct {
	Id               int64  `xorm:"pk autoincr BIGINT(10)"`
	GameLevel        string `xorm:"not null VARCHAR(5)"`
	UserId           int64  `xorm:"not null BIGINT(10)"`
	PlayTime         int64  `xorm:"not null BIGINT(10)"`
	GameScore        int64  `xorm:"not null BIGINT(10)"`
	GameRoomId       string `xorm:"not null VARCHAR(20)"`
	GameRewardSilver int64  `xorm:"not null BIGINT(10)"`
	CreateTime       int64  `xorm:"not null BIGINT(20)"`
	UpdateTime       int64  `xorm:"not null BIGINT(20)"`
}

func (p *LogGame) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "game_level":
		return p.GameLevel
	case "user_id":
		return p.UserId
	case "play_time":
		return p.PlayTime
	case "game_score":
		return p.GameScore
	case "game_room_id":
		return p.GameRoomId
	case "game_reward_silver":
		return p.GameRewardSilver
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	}
	return nil
}

func (p *LogGame) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type LogGames []*LogGame

func NewLogGames(cap int32) LogGames {
	return make(LogGames, 0, cap)
}

func (p LogGames) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetLogGame(cols []string, query string, args ...interface{}) (*LogGame, error) {
	obj := &LogGame{}
	ok, err := LogGameXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find LogGame by %s (%v)", query, args)
	}
	return obj, nil
}

func LogGameXorm() *xorm.Session {
	return orm.Xorm(1).Table("log_game")
}
