package gin

import (
	"fmt"
	"strconv"

	"github.com/Irfish/component/log"
	"github.com/gin-gonic/gin"
)

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
			context.Header("Access-Control-Allow-Origin", "*")
			if context.Request.Method == "OPTIONS" {
				//context.Header("Access-Control-Allow-Headers", "content-type,userId,token,tokenExpiredTime")
				context.Status(200)
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

