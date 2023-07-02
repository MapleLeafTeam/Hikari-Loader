package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://xovyafgv:YhY4V0awVyWccPJ7J8K30u2fbPH44tZM@tiny.db.elephantsql.com/xovyafgv"
	datebase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

type Product struct {  
    gorm.Model  
    Code string  
    Price uint  
}
db.AutoMigrate(&Product{})