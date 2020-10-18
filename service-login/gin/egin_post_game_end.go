package gin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
		}
		c.JSON(http.StatusOK, result)
	}()
	gameId := GetStringFromPostForm(c,"gameId")
	gameEnd := GetStringFromPostForm(c,"gameEnd")
	var g []GameResult
	err:= json.Unmarshal([]byte(gameEnd),g)
	if err!=nil{
		e = fmt.Errorf(err.Error())
		return
	}
	gameManager.GameEnd(gameId,g)
	result["state"] = true
}