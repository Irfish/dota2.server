package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EginGetReloadBlackList struct {
}

func NewEginGetReloadBlackList() EginGetReloadBlackList {
	p := EginGetReloadBlackList{}
	return p
}

func (p *EginGetReloadBlackList) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	LoadBlackList()
	result["state"] = true
}