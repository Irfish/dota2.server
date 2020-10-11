package gin

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/str"
	"github.com/Irfish/dota2.server/service-login/orm"
	"strconv"
	"time"
)

var gameManager = NewGameManager()

type GameManager struct {
	Games  map[int64]*Game
}

func NewGameManager() *GameManager {
	m := &GameManager{}
	m.Games = make(map[int64]*Game,0)
	return m
}
//创建游戏
func (m *GameManager)GameCreated() (bool,int64) {
	game := &Game{}
	gameID := genGameID()
	if gameID==0 {
		return false,0
	}
	game.GameID = gameID
	game.CreateTime = time.Now().Unix()
	game.State = GAME_STATE_IDLE
	game.Players = make(map[string]*Player,0)
	m.Games[gameID] = game
	return true,gameID
}

//玩家进入
func (m *GameManager)PlayerEnter(steamId string,gameId int64,index int) bool {
	game := m.Games[gameId]
	player :=&Player{
		SteamId:steamId,
	}
	exit,u :=GetUser(steamId)
	if exit {
		player.GameState = PLAYER_STATE_IN_GAME
		player.Silver = u.SteamSilver
		player.Gold = u.SteamGold
		player.VipExp = u.SteamVipExp
		player.Index = index
		player.UseTime = 5
		player.Items = GetItems(steamId)
	}else {
		return false
	}

	player.GameState = PLAYER_STATE_IDLE
	game.Players[steamId] =player
	return true
}

func (m *GameManager)GetGame(gameId int64) *Game {
	return  m.Games[gameId]
}

func (m *GameManager)GetPlayer(gameId int64,steamId string) *Player {
	game := m.GetGame(gameId)
	if game!=nil {
		return  game.Players[steamId]
	}
	return nil
}

func (m *GameManager)RefreshPlayer(gameId int64,steamId string)  {
	player:= m.GetPlayer(gameId,steamId)
	if player!=nil {
		exit,u :=GetUser(steamId)
		if exit {
			player.GameState = PLAYER_STATE_IN_GAME
			player.Silver = u.SteamSilver
			player.Gold = u.SteamGold
			player.VipExp = u.SteamVipExp
			player.Items = GetItems(steamId)
		}
	}
}

func (m *GameManager)GameEnd(gameId int64, steamId string, score int64,silver int64)  {
	exit,u :=GetUser(steamId)
	if !exit {
		return
	}
	g:= m.GetGame(gameId)
	if g!=nil {
		t := time.Now().Unix()
		u.SteamSilver = u.SteamSilver+ silver
		u.UpdateTime = t
		_,err:= orm.UserXorm().Where("steam_id=?",steamId).Update(&u)
		if err!=nil {
			log.Debug("%s",err.Error())
			return
		}
		logGame := orm.LogGame{}
		logGame.UpdateTime = t
		logGame.UserId = u.Id
		logGame.GameRoomId = gameId
		logGame.GameRewardSilver = silver
		logGame.GameScore = score
		logGame.GameLevel = ""
		logGame.PlayTime = 0
		logGame.CreateTime = t
		_,err = orm.LogGameXorm().Insert(&logGame)
		if err!=nil {
			log.Debug("%s",err.Error())
			return
		}
		g.State = GAME_STATE_END
		delete(m.Games,g.GameID)
	}
}

//生成游戏ID
func genGameID() int64 {
	str:= str.RandomNumber(6)
	if str ==""{
		log.Debug("RandomNumber 0")
		return 0
	}
	roomId,e:= strconv.ParseInt(str,10,64)
	if e!=nil {
		log.Debug("ParseInt err:",e.Error())
		return 0
	}
	return  roomId
}







