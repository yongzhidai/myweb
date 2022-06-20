package biz

import (
	"myweb/dto"
	"myweb/service"
)

func ListArticle(req dto.ListReq) ([]dto.ArticleListResp, error) {
	return service.ListArticle(req)
}

func SaveArticle(req dto.EditReq) error {
	return service.SaveArticle(req)
}
