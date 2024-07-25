package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPathParamInt(ctx *gin.Context, key string) (int, error) {
	idStr := ctx.Param(key)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
