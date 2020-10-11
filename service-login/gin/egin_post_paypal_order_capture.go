package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PayPalOrderCapture struct {

}

func NewPayPalOrderCapture() PayPalOrderCapture {
	p := PayPalOrderCapture{}
	return p
}

func (p *PayPalOrderCapture) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
}
