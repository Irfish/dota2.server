package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
)

var enterGameMutex *sync.RWMutex

type EnterGame struct {

}

func NewEnterGame() EnterGame {
	p := EnterGame{}
	enterGameMutex = new(sync.RWMutex)
	return p
}

func (p *EnterGame) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
			result["errCode"] = ERRORCODE_SERVER_ERR
		}
		c.JSON(http.StatusOK, result)
		enterGameMutex.Unlock()
	}()
	enterGameMutex.Lock()
	players := GetStringFromPostForm(c,"players")
	ids := strings.Split(players,",")
	playerIllegal := true
	for _,v :=range ids {
		if v!="0" {
			exit,_:= GetAndCreate(v)
			if !exit {
				playerIllegal = false
			}
		}
	}
	if playerIllegal {
		steamID := ids[0]
		inGame,gameId:= gameManager.CheckPlayerInGame(steamID)
		if !inGame {
			state,id:= gameManager.GameCreated()
			if state {
				gameId = id
				for index:=0;index< len(ids); index++ {
					steamId:= ids[index]
					if steamId!="0" {
						state := gameManager.PlayerEnter(ids[index],gameId,index)
						if !state {
							e = fmt.Errorf("player Enter game failed %d",gameId)
							return
						}
					}
				}
			}else {
				e = fmt.Errorf("create game failed")
				return
			}
		}
		game:= gameManager.GetGame(gameId)
		if game!=nil {
			p:= gameManager.GetPlayer(gameId,steamID)
			if p==nil{
				e = fmt.Errorf("player not exit steamid= %s",ids[0])
				return
			}
			result["player"] = p
			result["gameID"] = gameId
			return
		}else {
			e = fmt.Errorf("can not found game with gameid %s",gameId)
			return
		}
	}
	e = fmt.Errorf("game creat failed ")
	return
}





