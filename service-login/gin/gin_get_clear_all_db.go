package gin

import (
	"fmt"
	"github.com/Irfish/dota2.server/service-login/base"
	"github.com/Irfish/dota2.server/service-login/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EinGetClearAllDB struct {
}

func NewEinGetClearAllDB() EinGetClearAllDB {
	p := EinGetClearAllDB{}
	return p
}

func (p *EinGetClearAllDB) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	if base.Server.LogLevel!="debug" {
		return
	}
	{
		users := orm.NewUsers(0)
		err:= orm.UserXorm().Cols("id").Find(&users)
		if err!=nil{
			e = fmt.Errorf(err.Error())
			return
		}
		for _,u:=range users {
			orm.UserXorm().Where("id=?",u.Id).Delete(u)
		}
	}

	{
		userBags := orm.NewUserBags(0)
		err:= orm.UserBagXorm().Cols("id").Find(&userBags)
		if err!=nil{
			e = fmt.Errorf(err.Error())
			return
		}
		for _,m:=range userBags {
			orm.UserBagXorm().Where("id=?",m.Id).Delete(m)
		}
	}

	{
		logUseItems := orm.NewLogUseItems(0)
		err:= orm.LogUseItemXorm().Cols("id").Find(&logUseItems)
		if err!=nil{
			e = fmt.Errorf(err.Error())
			return
		}
		for _,m:=range logUseItems {
			orm.LogUseItemXorm().Where("id=?",m.Id).Delete(m)
		}
	}

	{
		logBuyGoodss := orm.NewLogBuyGoodss(0)
		err:= orm.LogBuyGoodsXorm().Cols("id").Find(logBuyGoodss)
		if err!=nil{
			e = fmt.Errorf(err.Error())
			return
		}
		for _,m:=range logBuyGoodss {
			orm.LogBuyGoodsXorm().Where("id=?",m.Id).Delete(m)
		}
	}
	gameManager.Games = make(map[string]*Game,0)
}