package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var insertSourceMutex *sync.RWMutex

type GenCardKey struct {
}

func NewGenCardKey() GenCardKey {
	insertSourceMutex = new(sync.RWMutex)
	p := GenCardKey{}
	return p
}

func (p *GenCardKey) handle(c *gin.Context) {
	insertSourceMutex.Lock()
	var e error
	result := gin.H{}
	defer func() {
		insertSourceMutex.Unlock()
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()

	money := StringToInt(c.Query("money"))
	value := StringToInt(c.Query("value"))
	count := StringToInt(c.Query("count"))
	if GenCardKeys(money,value,count) {
		if ExportTxtFile() {
			result["state"]=true
			return
		}else {
			e = fmt.Errorf("save keys to file failed")
			return
		}
	}else{
		e = fmt.Errorf("gen card key failed")
		return
	}
}