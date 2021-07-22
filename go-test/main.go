package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// _ "github.com/go-sql-driver/mysql"
	// // go import の影響で消える
	// _ "github.com/hrdtknr/GoAndMySQL"
)

func main() {

	// db := test.OpenDB(os.Getenv("DRIVER"), os.Getenv("DSN"))
	// defer test.CloseDB(db)

	// if err := db.Ping(); err != nil {
	// 	log.Fatal("db.Ping failed:", err)
	// }

	fmt.Println("start")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "hello, docker container\n")
	})

	// ローカルだと.env読めてない
	// コンテナ起動すればOK
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
