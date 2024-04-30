package old

import "github.com/gin-gonic/gin"

var oldList = make(map[string][]string)

func Init() {
	oldList = map[string][]string{}
}

func Set(ctx *gin.Context) {
	ctx.Request.ParseForm()
	oldList = ctx.Request.PostForm
}

func Get() map[string][]string {
	return oldList
}
