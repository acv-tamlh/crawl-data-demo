package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "postgres"
  password = ""
  dbname = "kikan_production"
)

func main()  {
  fmt.Println("main header")
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
    "password=%s dbname=%s sslmode=disabled", host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  fmt.Println("connected")
  defer db.Close()

}
