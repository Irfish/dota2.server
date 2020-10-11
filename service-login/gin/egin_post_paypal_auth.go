package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PayPalOauth struct {

}

func NewPayPalOauth() PayPalOauth {
	p := PayPalOauth{}
	return p
}

func (p *PayPalOauth) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	steamID, ok := c.GetPostForm("SteamID")
	if !ok {
		e = fmt.Errorf("the key steamID not found")
		return
	}
	//todo check steamId in Database
	if steamID=="" {
		e = fmt.Errorf("steamID is illegal")
		return
	}
	token,err:= payPalManager.PayPalOauth()
	if err!=nil{
		e = fmt.Errorf(err.Error())
		return
	}
	result["token"] = token
}