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
		items,limitItems:= GetItems(steamId)
		player.GameState = PLAYER_STATE_IN_GAME
		player.Silver = u.SteamSilver
		player.Gold = u.SteamGold
		player.VipExp = u.SteamVipExp
		player.Index = index
		player.Items = items
		player.LimitItems = limitItems
		player.ID = u.Id
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

func (m *GameManager)GetUser(steamId string) *Player {
	for _,g:=range m.Games {
		for _,p:=range g.Players {
			if p.SteamId == steamId {
				return p
			}
		}
	}
	return nil
}

func (m *GameManager)RefreshPlayer(gameId,steamId string)  {
	player:= m.GetPlayer(gameId,steamId)
	if player!=nil {
		exit,u :=GetUser(steamId)
		items,limitItems:= GetItems(steamId)
		if exit {
			player.GameState = PLAYER_STATE_IN_GAME
			player.Silver = u.SteamSilver
			player.Gold = u.SteamGold
			player.VipExp = u.SteamVipExp
			player.Items = items
			player.LimitItems = limitItems
			player.ID = u.Id
		}
	}
}

func GetExpByLevel(level int64) int64 {
	switch level {
	case 1:
		return 10
	case 2:
		return 20
	case 3:
		return 30
	case 4:
		return 40
	}
	return 0
}

func (m *GameManager)GameEnd(gameId string,result []GameResult,gameState,gameLevel int64)  {
	g:= m.GetGame(gameId)
	g.State = GAME_STATE_END
	if g!=nil {
		t := time.Now().Unix()
		playTime :=t-g.CreateTime
		gameRanks := make([]GameRank,0)
		exp:= GetExpByLevel(gameLevel)
		if gameState!=GAME_RESULT_WIN {
			exp = 0
		}
		for _,v:=range result {
			m.gameEndLog(gameId,v.SteamId,v.Score,v.Silver,playTime,t,exp)
			gameRanks = append(gameRanks,GameRank{	v.SteamId,v.Score,playTime})
		}
		if gameState==GAME_RESULT_WIN {
			gameRankManager.Update(gameRanks)
		}
	}
	delete(m.Games,g.GameID)
}

func (m *GameManager)gameEndLog(gameId , steamId string, score ,silver,playTime,t,exp int64)  {
	exit,u :=GetUser(steamId)
	if !exit {
		return
	}
	g:= m.GetGame(gameId)
	if g!=nil {
		u.SteamSilver = u.SteamSilver + silver
		u.UpdateTime = t
		u.SteamVipExp = u.SteamVipExp + exp
		_,err:= orm.UserXorm().Where("steam_id=?",steamId).Cols("steam_silver","steam_vip_exp","update_time").Update(&u)
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
	}
}

//生成游戏ID
func genGameID() string {
	return  str.RandomStringWithBigChar(10)
}

func CheckGameID(gameId string) (bool,int)  {
	g:= gameManager.GetGame(gameId)
	if g!=nil{
		return true, ERRORCODE_GAME_NOT_EXIT
	}
	return false,0
}

func CheckSteamID(gameId string,steamId string) (bool,int)  {
	p:=gameManager.GetPlayer(gameId,steamId )
	if p!=nil{
		return true, ERRORCODE_PLAYER_NOT_EXIT
	}
	return false,0
}






