package html

import "github.com/gin-gonic/gin"

func LoadHTML(r *gin.Engine) {

	//internal/modules/moduleName/html/view.tmpl
	r.LoadHTMLGlob("internal/**/**/**/*tmpl")

}
