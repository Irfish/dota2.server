package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type LogCardKey struct {
	Id         int64 `xorm:"pk autoincr BIGINT(10)"`
	CardKeyId  int64 `xorm:"not null BIGINT(10)"`
	UserId     int64 `xorm:"not null BIGINT(10)"`
	UseTime    int64 `xorm:"not null BIGINT(10)"`
	Cost       int64 `xorm:"not null BIGINT(10)"`
	CreateTime int64 `xorm:"not null BIGINT(20)"`
	UpdateTime int64 `xorm:"not null BIGINT(20)"`
}

func (p *LogCardKey) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "card_key_id":
		return p.CardKeyId
	case "user_id":
		return p.UserId
	case "use_time":
		return p.UseTime
	case "cost":
		return p.Cost
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	}
	return nil
}

func (p *LogCardKey) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type LogCardKeys []*LogCardKey

func NewLogCardKeys(cap int32) LogCardKeys {
	return make(LogCardKeys, 0, cap)
}

func (p LogCardKeys) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetLogCardKey(cols []string, query string, args ...interface{}) (*LogCardKey, error) {
	obj := &LogCardKey{}
	ok, err := LogCardKeyXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find LogCardKey by %s (%v)", query, args)
	}
	return obj, nil
}

func LogCardKeyXorm() *xorm.Session {
	return orm.Xorm(1).Table("log_card_key")
}
