package gin

import (
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
	steamId := GetStringFromPostForm(c,"steamId")
	gameId := GetInt64FromPostForm(c,"gameId")
	score := GetInt64FromPostForm(c,"score")
	silver := GetInt64FromPostForm(c,"silver")
	gameManager.GameEnd(gameId,steamId,score,silver)
	result["player"] = gameManager.GetPlayer(gameId,steamId)
}