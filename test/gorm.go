package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

func main() {
  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()

  // 自动迁移模式
  db.AutoMigrate(&Product{})

  // 创建
  db.Create(&Product{Code: "L1212", Price: 1000})
  var product1 = Product{Code: "99", Price:88}
  db.Create(&product1)

  // 读取
  var product Product
  //db.First(&product, 1) // 查询id为1的product
  db.First(&product, "Code = ?", "99") // 查询code为l1212的product
  fmt.Println(product.Code)
  fmt.Println(product.Price)
  // 更新 - 更新product的price为2000
  db.Model(&product).Update("Price", 2000)

  // 删除 - 删除product
  //db.Delete(&product)
}
