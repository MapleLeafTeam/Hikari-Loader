package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://xovyafgv:YhY4V0awVyWccPJ7J8K30u2fbPH44tZM@tiny.db.elephantsql.com/xovyafgv"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	new_anime := Anime{id: 1, url: "https://example.com/anime1.mp4"}

	db.Create(&new_anime) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
}

type Anime struct {
	id  int
	url string
}
