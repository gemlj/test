package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"

)


type Person struct {
 ID uint 
 FirstName string 
 LastName string 
}

func main() {
  db, err := gorm.Open("sqlite3", "gorm2.db")
  if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()
// 自动迁移模式
  db.AutoMigrate(&Person{})

 var person1 = Person{FirstName: "John", LastName: "Doe"}
 db.Create(&person1)
 //db.Create(&Product{Code: "L1212", Price: 1000})

 db.Create(&Person{FirstName: "LI",LastName: "JIN"})
 var p3 Person // identify a Person type for us to store the results in
 db.First(&p3) // Find the first record in the Database and store it in p3

 fmt.Println(p3.LastName) // print out our record from the database

}
