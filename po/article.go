package po

import (
	"myweb/db"
	"time"
)

type Article struct {
	ID         int64 `gorm:"primary_key"`
	Title      string
	Content    string `gorm:"size:500"`
	CreateTime time.Time
}

func SyncArticleTable() {
	db.DB.AutoMigrate(Article{})
}
