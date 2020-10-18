package gin

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/str"
	"github.com/Irfish/dota2.server/service-login/orm"
	"time"
)

var gameManager = NewGameManager()

type GameManager struct {
	Games  map[string]*Game
}

func NewGameManager() *GameManager {
	m := &GameManager{}
	m.Games = make(map[string]*Game,0)
	return m
}
//创建游戏
func (m *GameManager)GameCreated() (bool,string) {
	game := &Game{}
	gameID := genGameID()
	game.GameID = gameID
	game.CreateTime = time.Now().Unix()
	game.State = GAME_STATE_IDLE
	game.Players = make(map[string]*Player,0)
	m.Games[gameID] = game
	return true,gameID
}

func (m *GameManager)CheckPlayerInGame(steamId string ) (bool,string) {
	for gameId,game :=range m.Games {
		players := game.Players
		if _,b:= players[steamId];b{
			return true,gameId
		}
	}
	return false,""
}

//玩家进入
func (m *GameManager)PlayerEnter(steamId ,gameId string,index int) bool {
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

func (m *GameManager)GetGame(gameId string) *Game {
	return  m.Games[gameId]
}

func (m *GameManager)GetPlayer(gameId string,steamId string) *Player {
	game := m.GetGame(gameId)
	if game!=nil {
		return  game.Players[steamId]
	}
	return nil
}

func (m *GameManager)RefreshPlayer(gameId,steamId string)  {
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

func (m *GameManager)GameEnd(gameId string,result []GameResult)  {
	g:= m.GetGame(gameId)
	if g!=nil {
		t := time.Now().Unix()
		playTime :=t-g.CreateTime
		gameRanks := make([]GameRank,0)
		for _,v:=range result {
			m.gameEndLog(gameId,v.SteamId,v.Score,v.Silver,playTime,t)
			gameRanks = append(gameRanks,GameRank{	v.SteamId,v.Score,playTime})
		}
		gameRankManager.Update(gameRanks)
	}
}

func (m *GameManager)gameEndLog(gameId , steamId string, score ,silver,playTime,t int64)  {
	exit,u :=GetUser(steamId)
	if !exit {
		return
	}
	g:= m.GetGame(gameId)
	if g!=nil {
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
		logGame.PlayTime = playTime
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
func genGameID() string {
	return  str.RandomString(6)
}







