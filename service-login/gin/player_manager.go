package gin

import (
	"fmt"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/xorm"
	"github.com/Irfish/dota2.server/service-login/orm"
	"time"
)

type PlayerManager struct {
	Players map[string]*Player
}


func GetAndCreate(id string) (bool,orm.User ) {
	exit,u := GetUser(id)
	if exit {
		return  true,u
	}else {
		state,u:= NewUser(id)
		if state {
			return true,u
		}else {
			log.Debug("insert user failed:")
			return false,orm.User{}
		}
	}
	log.Debug("GetAndCreate user failed:")
	return false,orm.User{}
}

func GetUser(steamId string)  (bool,orm.User ) {
	u := orm.User{}
	exist, err := orm.UserXorm().Where("steam_id=?",steamId).Get(&u)
	if err !=nil {
		log.Debug("do get user err:%s",err.Error())
		return false,orm.User{}
	}
	if exist {
		return true,u
	}
	return false,orm.User{}
}

func NewUser(steamId string)  (bool,orm.User ) {
	t := time.Now().Unix()
	u := orm.User{
		SteamId: steamId,
		SteamGold:0,
		SteamSilver:0,
		SteamVipExp:0,
		CreateTime: t,
		UpdateTime: t,
	}
	affected, err := orm.UserXorm().Insert(u)
	if err !=nil {
		log.Debug("insert user err:%s",err.Error())
	}
	if affected>0 {
		return true,u
	}
	return false,orm.User{}
}

func GetItems(steamId string) []Item {
	exit,u := GetUser(steamId)
	if exit {
		userID:= u.Id
		userBags := []orm.UserBag{}
        count,err:=orm.UserBagXorm().Where("user_id=? and item_state=?",userID,ITEM_STATE_HAVE).FindAndCount(&userBags)
		if err !=nil {
			log.Debug("insert user err:%s",err.Error())
		}
		if count>0 {
			items := make([]Item,count)
			for i,item := range userBags{
				items[i] = Item{item.ItemId,item.ItemCount}
			}
			return items
		}
	}
	return  []Item{}
}

func PlayerBuyItem(steamId string,itemId int64,cost int64,count int64) bool {
	t := time.Now().Unix()
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false
	}
	if u.SteamGold < cost {
		log.Debug("gold not enough")
		return false
	}
	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}

	momey := u.SteamGold-cost
	user:= orm.User{
		SteamGold:momey,
		UpdateTime: t,
	}
	//更新玩家金币
	{
		s1:= s.Table("user")
		effect,err:= s1.ID(u.Id).Update(&user)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if effect!=1 {
			e = fmt.Errorf("update gold failed in user table")
			return false
		}
	}
	//将道具存入背包
	{
		s2:= s.Table("user_bag")
		ub := orm.UserBag{}
		b, err:= s2.Where("user_id=? and item_id=?",u.Id,itemId).Get(&ub)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if b {
			if ub.ItemState == ITEM_STATE_HAVE {
				ub.ItemCount = ub.ItemCount+count
			}else {
				ub.ItemState = ITEM_STATE_HAVE
				ub.ItemCount = count
			}
			ub.UpdateTime = t
			effectNum, err:= s2.Where("user_id=? and item_id=?",u.Id,itemId).Update(&ub)
			if err !=nil {
				e = fmt.Errorf(err.Error())
				return false
			}
			if effectNum!=1 {
				e = fmt.Errorf("update buy itme failed in userbag table")
				return false
			}
		}else {
			userBag:= orm.UserBag{
				UserId: u.Id,
				ItemId: itemId,
				ItemCount: count,
				ItemState: ITEM_STATE_HAVE,
				CreateTime: t,
				UpdateTime: t,
			}
			effect,err:= s2.Insert(&userBag)
			if err !=nil {
				e = fmt.Errorf(err.Error())
				return false
			}
			if effect!=1 {
				e = fmt.Errorf("insert buy itme failed in userbag table")
				return false
			}
		}
		s.Commit()
	}
	//记录使用金币兑换道具日志
	{
		logUseItem := orm.LogUseItem{
			CreateTime: t,
			UpdateTime: t,
			ItemId: itemId,
			UserId: u.Id,
			ItemUseState:ITEM_FROM_BUY,
		}
		_,err:=orm.LogUseItemXorm().Insert(&logUseItem)
		if err !=nil {
			log.Debug("%s",err.Error())
		}
	}
	return true
}


func PlayerUseItem(steamId string,itemId int64,count int64)  {

}

func CheckCardKey(code string) (bool,int) {
	card := orm.CardKey{}
	exit,err:=orm.CardKeyXorm().Where("key_code=?",code).And("key_state=?",KEY_STATE_NORMAL) .Get(&card)
	if err!=nil {
		log.Debug(err.Error())
		return false,ERRORCODE_CARD_KEY_USED
	}
	if !exit {
		log.Debug("code(%s)  not exit",code)
		return false,ERRORCODE_CARD_KEY_USED
	}
	return true,0
}

func PlayerUseCardKey(steamId string,code string) (e error ) {
	card := orm.CardKey{}
	exit,err:=orm.CardKeyXorm().Where("key_code=?",code).And("key_state=?",KEY_STATE_NORMAL) .Get(&card)
	if err!=nil {
		e= fmt.Errorf(err.Error())
	}
	if !exit {
		e=fmt.Errorf("code(%s)  not exit",code)
	}
	exit, u := GetUser(steamId)
	if !exit {
		e= fmt.Errorf("steamId(%s) not exit",steamId)
		return
	}
	t := time.Now().Unix()

	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}
		s.Close()

	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return
	}
	s1:= s.Table("card_key")

	//更新code状态
	{
		card.KeyState = KEY_STATE_USED
		card.UpdateTime= t
		effect,err:= s1.Where("key_code=?",code).Update(&card)
		if err!=nil {
			e = fmt.Errorf("%s",err.Error())
			return
		}
		if effect !=1 {
			e = fmt.Errorf("update key_code state failed")
			return
		}
	}
	//更新玩家金币
	{
		s2:= s.Table("user")
		u.UpdateTime = t
		u.SteamGold = u.SteamGold+card.KeyValue
		effect,err:= s2.Where("id=?",u.Id).Update(&u)
		if err!=nil {
			e = fmt.Errorf("%s",err.Error())
			return
		}
		if effect !=1 {
			e = fmt.Errorf("update steam_gold state failed")
			return
		}
	}
	s.Commit()

	//记录code使用日志
	{
		log:= orm.LogCardKey{}
		log.UpdateTime = t
		log.CreateTime = t
		log.UserId = u.Id
		log.CardKeyId= card.Id
		log.Cost = card.KeyRmb
		log.UseTime = 0
		effect,err:= orm.LogCardKeyXorm().Insert(&log)
		if err !=nil {
			e = fmt.Errorf("%s",err.Error())
			return
		}
		if effect !=1 {
			e = fmt.Errorf("insert  card key log failed")
			return
		}
	}
	return nil
}

func UpdateGold(steamId string,cost int64) bool {
	t := time.Now().Unix()
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false
	}
	if u.SteamGold < cost {
		log.Debug("gold not enough")
		return false
	}
	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}else {
			s.Commit()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}

	momey := u.SteamGold-cost
	user:= orm.User{
		SteamGold:momey,
		UpdateTime: t,
	}
	//更新玩家金币
	{
		s1:= s.Table("user")
		effect,err:= s1.ID(u.Id).Update(&user)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if effect!=1 {
			e = fmt.Errorf("update gold failed in user table")
			return false
		}
	}
	return true
}


func UpdateSilver(steamId string,cost int64) bool {
	t := time.Now().Unix()
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false
	}
	if u.SteamSilver < cost {
		log.Debug("silver not enough")
		return false
	}
	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}else {
			s.Commit()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}

	momey := u.SteamSilver-cost
	user:= orm.User{
		SteamSilver:momey,
		UpdateTime: t,
	}
	//更新玩家金币
	{
		s1:= s.Table("user")
		effect,err:= s1.ID(u.Id).Update(&user)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if effect!=1 {
			e = fmt.Errorf("update silver failed in user table")
			return false
		}
	}
	return true
}

func UpdateExp(steamId string,num int64) bool {
	t := time.Now().Unix()
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false
	}
	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}else {
			s.Commit()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}

	momey := u.SteamVipExp+num
	user:= orm.User{
		SteamVipExp:momey,
		UpdateTime: t,
	}
	//更新玩家金币
	{
		s1:= s.Table("user")
		effect,err:= s1.ID(u.Id).Update(&user)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if effect!=1 {
			e = fmt.Errorf("update vip exp failed in user table")
			return false
		}
	}
	return true
}

func UpdateItem(steamId string, itemId int64, count int64) bool {
	t := time.Now().Unix()
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false
	}
	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}else {
			s.Commit()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}

	//将道具存入背包
	{
		s2:= s.Table("user_bag")
		ub := orm.UserBag{}
		b, err:= s2.Where("user_id=? and item_id=?",u.Id,itemId).Get(&ub)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if b {
			if ub.ItemState == ITEM_STATE_HAVE {
				ub.ItemCount = ub.ItemCount+count
			}else {
				ub.ItemState = ITEM_STATE_HAVE
				ub.ItemCount = count
			}
			ub.UpdateTime = t
			effectNum, err:= s2.Where("user_id=? and item_id=?",u.Id,itemId).Update(&ub)
			if err !=nil {
				e = fmt.Errorf(err.Error())
				return false
			}
			if effectNum!=1 {
				e = fmt.Errorf("update Update itme failed in userbag table")
				return false
			}
		}else {
			userBag:= orm.UserBag{
				UserId: u.Id,
				ItemId: itemId,
				ItemCount: count,
				ItemState: ITEM_STATE_HAVE,
				CreateTime: t,
				UpdateTime: t,
			}
			effect,err:= s2.Insert(&userBag)
			if err !=nil {
				e = fmt.Errorf(err.Error())
				return false
			}
			if effect!=1 {
				e = fmt.Errorf("insert Update itme failed in userbag table")
				return false
			}
		}
		s.Commit()
	}
	return true
}

func CheckGold(steamId string, cost int64) (bool,int) {
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false,ERRORCODE_PLAYER_NOT_EXIT
	}
	if u.SteamGold < cost {
		log.Debug("gold not enough")
		return false,ERRORCODE_GOLD_NOT_ENOUGH
	}
	return true,0
}

func CheckSilver(steamId string, cost int64) (bool,int) {
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false,ERRORCODE_PLAYER_NOT_EXIT
	}
	if u.SteamSilver < cost {
		log.Debug("silver not enough")
		return false,ERRORCODE_SILVER_NOT_ENOUGH
	}
	return true,0
}

