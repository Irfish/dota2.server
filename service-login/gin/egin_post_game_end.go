package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type GameEnd struct {
}

func NewGameEnd() GameEnd {
	p := GameEnd{}
	return p
}

func (p *GameEnd) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
			result["errCode"] = ERRORCODE_SERVER_ERR
		}
		c.JSON(http.StatusOK, result)
	}()
	gameId := GetStringFromPostForm(c,"gameId")
	if b,i := CheckGameID(gameId);!b {
		result["errCode"] = i
		return
	}
	gameState := GetInt64FromPostForm(c,"gameState")
	gameLevel := GetInt64FromPostForm(c,"gameLevel")
	player := GetStringFromPostForm(c,"player")
	Score := GetStringFromPostForm(c,"score")
	Sliver := GetStringFromPostForm(c,"silver")
	g := make([]GameResult,0)
	ids := strings.Split(player,",")
	scores := strings.Split(Score,",")
	slivers := strings.Split(Sliver,",")
	for i:=0;i< len(ids);i++ {
		steam:= ids[i]
		if steam!="0" {
			sc, err := strconv.ParseInt(scores[i],10,64)
			if err != nil {
				e=fmt.Errorf("%s can not convert to int: %s",scores[i],err.Error())
				return
			}
			sl, err := strconv.ParseInt(slivers[i],10,64)
			if err != nil {
				e=fmt.Errorf("%s can not convert to int: %s",slivers[i],err.Error())
				return
			}
			s := GameResult{
				SteamId: steam,
				GameID:gameId,
				Score: sc,
				Silver: sl,
			}
			g=append(g,s)
		}
	}
	gameManager.GameEnd(gameId,g,gameState,gameLevel)
	result["state"] = true
}