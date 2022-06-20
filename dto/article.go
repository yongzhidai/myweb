package dto

import (
	"errors"
	"myweb/po"
)

type EditReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (e *EditReq) CheckParam() error {
	if e.Title == "" {
		return errors.New("文章标题不能为空")
	}
	if e.Content == "" {
		return errors.New("文章内容不能为空")
	}
	return nil
}

type ListReq struct {
	Offset   int32 `json:"offset"`
	PageSize int32 `json:"pageSize"`
}

func (l *ListReq) CheckParam() error {
	if l.Offset < 0 {
		return errors.New("偏移量不合法")
	}
	if l.PageSize <= 0 || l.PageSize > 50 {
		return errors.New("页大小不支持")
	}
	return nil
}

type ArticleListResp struct {
	Id         int64
	Title      string
	Content    string
	CreateTime string
}

func BuildArticleListResp(articles []po.Article) []ArticleListResp {
	var result []ArticleListResp
	for _, article := range articles {
		resp := ArticleListResp{
			Id:         article.ID,
			Title:      article.Title,
			Content:    article.Content,
			CreateTime: article.CreateTime.Format("2006-01-02"),
		}
		result = append(result, resp)
	}
	return result
}
