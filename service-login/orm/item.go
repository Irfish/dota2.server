package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type Item struct {
	Id          int64 `xorm:"pk autoincr BIGINT(20)"`
	ItemType    int   `xorm:"not null INT(11)"`
	ItemUseType int   `xorm:"not null INT(11)"`
	UseTime     int64 `xorm:"not null BIGINT(20)"`
	CreateTime  int64 `xorm:"not null BIGINT(20)"`
	UpdateTime  int64 `xorm:"not null BIGINT(20)"`
}

func (p *Item) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "item_type":
		return p.ItemType
	case "item_use_type":
		return p.ItemUseType
	case "use_time":
		return p.UseTime
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	}
	return nil
}

func (p *Item) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type Items []*Item

func NewItems(cap int32) Items {
	return make(Items, 0, cap)
}

func (p Items) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetItem(cols []string, query string, args ...interface{}) (*Item, error) {
	obj := &Item{}
	ok, err := ItemXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find Item by %s (%v)", query, args)
	}
	return obj, nil
}

func ItemXorm() *xorm.Session {
	return orm.Xorm(0).Table("item")
}
