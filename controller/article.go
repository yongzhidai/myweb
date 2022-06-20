package controller

import (
	"github.com/gin-gonic/gin"
	"myweb/biz"
	"myweb/dto"
	"myweb/result"
	"net/http"
)

func ListArticle(c *gin.Context) {
	defer HandlePanic(c)
	//参数
	var req = new(dto.ListReq)
	c.BindJSON(req)
	err := req.CheckParam()
	if err != nil {
		c.JSON(http.StatusOK, result.ParamError(err.Error()))
		return
	}
	//业务流程
	resp, err := biz.ListArticle(*req)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, result.Success(resp))
}

func SaveArticle(c *gin.Context) {
	defer HandlePanic(c)
	//参数
	var req dto.EditReq
	c.BindJSON(&req)
	err := req.CheckParam()
	if err != nil {
		c.JSON(http.StatusOK, result.ParamError(err.Error()))
		return
	}
	//业务流程
	err = biz.SaveArticle(req)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, result.Success(nil))
}
