package handler

import (
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context)  {
	key := ctx.Query("key")
	if key == "" {
		msg := "params key should not empty"
		ctx.Render(Response(ctx, StatusCodeInternalServerError, msg, map[string]interface{}{}))
		return
	}

	res := make(map[string]interface{})
	res["key"] = key
	res["name"] = "the name of key"
	ctx.Render(Response(ctx, StatusCodeSuccess, "", res))
}