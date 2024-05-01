package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

)

func Set(ctx *gin.Context, key string, value string) {
	session := sessions.Default(ctx)

	session.Set(key, value)
	session.Save()
}

func Flash(ctx *gin.Context, key string) string {
	session := sessions.Default(ctx)

	response := session.Get(key)
	session.Save()

	session.Delete(key)
	session.Save()

	if response != nil {
		return response.(string)
	}

	return ""
}

func Get(ctx *gin.Context, key string) string {
	session := sessions.Default(ctx)

	response := session.Get(key)
	session.Save()

	if response != nil {
		return response.(string)
	}

	return ""
}

func Remove(ctx *gin.Context, key string) {
	session := sessions.Default(ctx)

	session.Delete(key)
	session.Save()
}
