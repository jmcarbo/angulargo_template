package main

import (
  "fmt"
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "time"
)

type User struct {
  Id int64
  Name string
  //Surname string
  //Email string
  Age int
  Birthday time.Time
}

func main(){
  var db gorm.DB
  var err error
  db, err = gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Connected")
  //db.CreateTable(User{})
  //db.AutoMigrate(User{})
  //user := User{Name: "jinzhu almeina", Age: 18, Birthday: time.Now()}
  //db.Save(&user)
  var users []User
  var user User
  db.Find(&users) 
  //fmt.Printf("%#v", users)
  db.Where(&User{Name: "jinzhu"}).First(&user)
  fmt.Printf("%#v", user)

}
