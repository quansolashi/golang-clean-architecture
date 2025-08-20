package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryInt64(ctx *gin.Context, query string, defaultValue int64) (int64, error) {
	str := strconv.FormatInt(defaultValue, 10)
	return strconv.ParseInt(ctx.DefaultQuery(query, str), 10, 64)
}

func GetParamUint64(ctx *gin.Context, param string, defaultValue uint64) (uint64, error) {
	return strconv.ParseUint(ctx.Param(param), 10, 64)
}
