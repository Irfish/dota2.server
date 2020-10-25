package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type UserBag struct {
	Id         int64 `xorm:"pk autoincr BIGINT(10)"`
	UserId     int64 `xorm:"not null BIGINT(20)"`
	ItemId     int64 `xorm:"not null BIGINT(20)"`
	ItemCount  int64 `xorm:"not null BIGINT(10)"`
	ItemState  int   `xorm:"not null INT(11)"`
	CreateTime int64 `xorm:"not null BIGINT(20)"`
	UpdateTime int64 `xorm:"not null BIGINT(20)"`
	UseState   int   `xorm:"not null INT(11)"`
	UseTime    int64 `xorm:"not null BIGINT(20)"`
}

func (p *UserBag) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "user_id":
		return p.UserId
	case "item_id":
		return p.ItemId
	case "item_count":
		return p.ItemCount
	case "item_state":
		return p.ItemState
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	case "use_state":
		return p.UseState
	case "use_time":
		return p.UseTime
	}
	return nil
}

func (p *UserBag) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type UserBags []*UserBag

func NewUserBags(cap int32) UserBags {
	return make(UserBags, 0, cap)
}

func (p UserBags) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetUserBag(cols []string, query string, args ...interface{}) (*UserBag, error) {
	obj := &UserBag{}
	ok, err := UserBagXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find UserBag by %s (%v)", query, args)
	}
	return obj, nil
}

func UserBagXorm() *xorm.Session {
	return orm.Xorm(0).Table("user_bag")
}
