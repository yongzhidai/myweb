package service

import (
	"myweb/db"
	"myweb/dto"
	"myweb/po"
	"time"
)

func ListArticle(req dto.ListReq) ([]dto.ArticleListResp, error) {
	var articles []po.Article
	db := db.DB.Limit(req.PageSize).Offset(req.Offset).Order("create_time desc").Find(&articles)
	if db.Error != nil {
		return nil, db.Error
	}
	return dto.BuildArticleListResp(articles), nil
}

func SaveArticle(req dto.EditReq) error {
	result := db.DB.Create(&po.Article{Title: req.Title, Content: req.Content, CreateTime: time.Now()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
