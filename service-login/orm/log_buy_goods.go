package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type LogBuyGoods struct {
	Id         int64 `xorm:"pk autoincr BIGINT(20)"`
	UserId     int64 `xorm:"not null BIGINT(20)"`
	ItemId     int64 `xorm:"not null BIGINT(20)"`
	CostMoney  int64 `xorm:"not null BIGINT(10)"`
	From       int   `xorm:"not null INT(11)"`
	Count      int64 `xorm:"not null BIGINT(10)"`
	CreateTime int64 `xorm:"not null BIGINT(20)"`
	UpdateTime int64 `xorm:"not null BIGINT(20)"`
}

func (p *LogBuyGoods) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "user_id":
		return p.UserId
	case "item_id":
		return p.ItemId
	case "cost_money":
		return p.CostMoney
	case "from":
		return p.From
	case "count":
		return p.Count
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	}
	return nil
}

func (p *LogBuyGoods) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type LogBuyGoodss []*LogBuyGoods

func NewLogBuyGoodss(cap int32) LogBuyGoodss {
	return make(LogBuyGoodss, 0, cap)
}

func (p LogBuyGoodss) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetLogBuyGoods(cols []string, query string, args ...interface{}) (*LogBuyGoods, error) {
	obj := &LogBuyGoods{}
	ok, err := LogBuyGoodsXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find LogBuyGoods by %s (%v)", query, args)
	}
	return obj, nil
}

func LogBuyGoodsXorm() *xorm.Session {
	return orm.Xorm(1).Table("log_buy_goods")
}
