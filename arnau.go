package main

import (
  "fmt"
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "time"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

type User struct {
  Id int64
  Login string
  Password string
  Company_id int64
  //Surname string
  //Email string
}

func (u User)TableName() string {
  return "res_users"
}

type Partner struct {
  Name string
}

func (p Partner)TableName() string {
  return "res_partner"
}

type Invoice struct {
  Number string
  Type string
  State string
  Date_due time.Time
  DateInvoice time.Time
  Partners []Partner
  PartnerId int64
  AmountTotal float64
  AmountTax float64
}

func (i Invoice)TableName() string {
  return "account_invoice"
}

type PurchaseOrder struct {
  Id          int64
  CreateUid   int64  
  CreateDate  time.Time
  WriteDate   time.Time
  WriteUid    int64
  Origin      string
  JournalId   int64
  DateOrder   time.Time
  PartnerId   int64
  DestAddressId  int64
  FiscalPosition  int
  AmountUntaxed  float64
  LocationId    int64
  CompanyId     int64
  AmountTax     float64
  State        string 
  PricelistId  int64
  WarehouseId  int64
  PaymentTermId int64
  PartnerRef    string
  DateApprove  time.Time
  AmountTotal  float64
  Name          string
  Notes        string
  InvoiceMethod  string
  Shipped       bool
  Validator     int64
  MinimumPlannedDate time.Time
  RequisitionId int64
}


func (p PurchaseOrder)TableName() string {
  return "purchase_order"
}

type Product struct {
  NameTemplate string
}

func (p Product)TableName() string {
  return "product_product"
}

func main(){
  var db gorm.DB
  var err error
  db, err = gorm.Open("postgres", "user=gorm dbname=realstate sslmode=disable")
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
  fmt.Printf("%#v", users)
  db.Where(&User{Login: "jmcarbo"}).First(&user)

  fmt.Printf("%#v", user)

  var partners []Partner
  db.Find(&partners) 
  //fmt.Printf("%#v", partners)
  for _,p := range partners {
    fmt.Println(p.Name)
  }
  var invoices []Invoice
  db.Where(&Invoice{Type: "out_invoice"}).Find(&invoices) 
  //fmt.Printf("%#v", invoices)
  var partner Partner
  for _,i := range invoices {
    db.Model(&i).Related(&partner, "PartnerId")
    fmt.Println(i.DateInvoice)
    fmt.Println(i.Number)
    fmt.Println(partner.Name)
    fmt.Println(i.AmountTotal)
  }

  var orders []PurchaseOrder
  db.Find(&orders)
  fmt.Printf("%#v", orders)
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
        Charset: "UTF-8",
  }))

  // This is set the Content-Type to "text/html; charset=ISO-8859-1"
  m.Get("/", func(r render.Render) {
    r.HTML(200, "hello", "world")
  })

  // This is set the Content-Type to "application/json; charset=ISO-8859-1"
  m.Get("/api", func(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })

  m.Get("/invoices", func(r render.Render) {
    r.JSON(200, invoices)
  })

  m.Get("/partners", func(r render.Render) {
    r.JSON(200, partners)
  })

  m.Run()
}
