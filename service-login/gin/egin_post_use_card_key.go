package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UseCardKey struct {
}

func NewUseCardKey() UseCardKey {
	p := UseCardKey{}
	return p
}

func (p *UseCardKey) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	steamID := GetStringFromPostForm(c,"steamId")
	gameID := GetInt64FromPostForm(c,"gameID")

	code := GetStringFromPostForm(c,"code")
	if PlayerUseCardKey(steamID,code) {
		gameManager.RefreshPlayer(gameID,steamID)
		p:= gameManager.GetPlayer(gameID,steamID)
		result["player"] = p
	}else {
		e = fmt.Errorf("use card code failed")
	}
}