package main

import(
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(w)
	fmt.Println("hello docker")
}

func main(){
	fmt.Println("start")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
	fmt.Println("end")
}


