package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type LogUseItem struct {
	Id           int64 `xorm:"pk autoincr BIGINT(20)"`
	UserId       int64 `xorm:"not null BIGINT(20)"`
	ItemId       int64 `xorm:"not null BIGINT(20)"`
	ItemUseState int   `xorm:"not null INT(11)"`
	CreateTime   int64 `xorm:"not null BIGINT(20)"`
	UpdateTime   int64 `xorm:"not null BIGINT(20)"`
}

func (p *LogUseItem) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "user_id":
		return p.UserId
	case "item_id":
		return p.ItemId
	case "item_use_state":
		return p.ItemUseState
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	}
	return nil
}

func (p *LogUseItem) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type LogUseItems []*LogUseItem

func NewLogUseItems(cap int32) LogUseItems {
	return make(LogUseItems, 0, cap)
}

func (p LogUseItems) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetLogUseItem(cols []string, query string, args ...interface{}) (*LogUseItem, error) {
	obj := &LogUseItem{}
	ok, err := LogUseItemXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find LogUseItem by %s (%v)", query, args)
	}
	return obj, nil
}

func LogUseItemXorm() *xorm.Session {
	return orm.Xorm(1).Table("log_use_item")
}
