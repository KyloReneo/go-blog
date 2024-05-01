package view

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/kyloReneo/go-blog/internal/modules/user/helpers"
	"github.com/kyloReneo/go-blog/pkg/converters"
	"github.com/kyloReneo/go-blog/pkg/sessions"

)

func WithGlobalData(ctx *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(ctx, "errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(ctx, "old"))
	data["AUTH"] = helpers.Auth(ctx)
	return data
}
