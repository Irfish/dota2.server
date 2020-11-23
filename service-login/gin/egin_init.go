package gin

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"encoding/json"
	"github.com/Irfish/component/log"
	"github.com/gin-gonic/gin"
)

var blackList []string

func init()  {
	LoadBlackList()
}

func LoadBlackList()  {
	data, err := ioutil.ReadFile("config/black_list.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &blackList)
	if err != nil {
		log.Fatal("%v", err)
	}
}

func GetIP(addr string) int64 {
	s := strings.Split(addr,".")
	numStr := s[len(s)-1]
	i, err := strconv.ParseInt(numStr,10,64)
	if err != nil {
		log.Debug("%s can not convert to int: %s",numStr,err.Error())
	}else {
		return i
	}
	return 0
}

func (s *Gin) GinInit(egin *gin.Engine) {
	egin.Use(
		func(context *gin.Context) {
			var e error
			defer func() {
				if e != nil {
					fmt.Println(e)
					context.JSON(200, gin.H{
						"debug": e.Error(),
					})
					context.Abort()
				}
			}()
			isBlack :=false
			for _,v:=range  blackList {
				if v==context.ClientIP() {
					isBlack =true
				}
			}
			if isBlack {
				return
			}
			context.Header("Access-Control-Allow-Origin", "*")
			if context.Request.Method == "OPTIONS" {
				context.Header("Access-Control-Allow-Headers", "content-type,userId,token,tokenExpiredTime")
				context.Status(200)
				return
			}
			k := context.GetHeader("Server-Key")
			if k=="" {
				e = fmt.Errorf("Server-Key is nil")
				return
			}
		},
	)
}

func StringToInt(s string) (i int)  {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Debug("StringToInt err:",err.Error())
		return
	}
	return
}

func GetStringFromPostForm(c *gin.Context,key string)  (v string) {
	v, ok := c.GetPostForm(key)
	if !ok {
		log.Debug("can not found key: %s",key)
		return
	}
	return
}

func GetInt64FromPostForm(c *gin.Context,key string) (i int64) {
	v, ok := c.GetPostForm(key)
	if !ok {
		log.Debug("can not found key: %s",key)
		return
	}
	i, err := strconv.ParseInt(v,10,64)
	if err != nil {
		log.Debug("%s can not convert to int: %s",key,err.Error())
		return
	}
	return
}

