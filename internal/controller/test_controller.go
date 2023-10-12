package controller

import (
	"cotton-gin-test/api/v1/test"
	"cotton-gin-test/common/response"
	"cotton-gin-test/internal/service"
	"github.com/gin-gonic/gin"
)

func Sum(ctx *gin.Context) {
	var req test.SumReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, err)
		return
	}
	data, err := service.Sum(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, data)
}
