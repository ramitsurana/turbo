
package main

import (
    "fmt"
    "net/http"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/qor/qor"
    "github.com/qor/admin"
)

// Create a GORM-backend model
type User struct {
  gorm.Model
  Name string
}

// Create another GORM-backend model
type Product struct {
  gorm.Model
  Name        string
  Description string
}

func main() {  

  DB, _ := gorm.Open("sqlite3", "turbo.db")
  DB.AutoMigrate(&User{}, &Product{})

  // Initalize
  Admin := admin.New(&qor.Config{DB: DB})

  Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})
  Admin.AddMenu(&admin.Menu{Name: "Monitoring", Link: "/monitor"})
  Admin.AddMenu(&admin.Menu{Name: "Tools", Link: "/tools"})

  // Create resources from GORM-backend model
  Admin.AddResource(&User{})
  Admin.AddResource(&Product{})
  //Admin.AddResource(&Monitoring{})  

  // Register route
  mux := http.NewServeMux()
  // amount to /admin, so visit `/admin` to view the admin interface
  Admin.MountTo("/admin", mux)

  Admin.SetSiteName("Turbo")

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}

