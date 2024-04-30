package old

import "github.com/gin-gonic/gin"

var oldList = make(map[string][]string)

func init() {
	oldList = map[string][]string{}
}

func set(ctx *gin.Context) {
	ctx.Request.ParseForm()
	oldList = ctx.Request.PostForm
}
