package html

import (
	"github.com/gin-gonic/gin"

	"github.com/kyloReneo/go-blog/internal/providers/view"
)

func Render(c *gin.Context, code int, path string, data gin.H) {

	data = view.WithGlobalData(c, data)
	c.HTML(code, path, data)
}
