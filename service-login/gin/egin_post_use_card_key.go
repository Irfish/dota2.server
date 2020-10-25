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
			result["errCode"] = ERRORCODE_SERVER_ERR
		}
		c.JSON(http.StatusOK, result)
	}()

	gameID := GetStringFromPostForm(c,"gameID")
	if b,i := CheckGameID(gameID);!b {
		result["errCode"] = i
		return
	}

	steamID := GetStringFromPostForm(c,"steamId")
	if b,i := CheckSteamID(gameID,steamID);!b {
		result["errCode"] = i
		return
	}

	code := GetStringFromPostForm(c,"code")
	if b,i := CheckCardKey(code);!b {
		result["errCode"] = i
		return
	}

	err:= PlayerUseCardKey(steamID,code)
	if err==nil  {
		gameManager.RefreshPlayer(gameID,steamID)
		p:= gameManager.GetPlayer(gameID,steamID)
		result["player"] = p
	}else {
		e = fmt.Errorf(err.Error())
	}
}