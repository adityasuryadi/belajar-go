package main

import (
	"log"
	"net/http"
)

func main() {
	// i:= 21
	// j:= true
	// fmt.Printf("%T \n",i)
	// fmt.Println("%")
	// fmt.Println(j)
	// j = false
	// fmt.Println(j)
	// fmt.Printf("%U \n",'Я')
	// fmt.Printf("%d \n",10)
	// fmt.Printf("%o \n",25)
	// fmt.Printf("%x \n","f")
	// fmt.Printf("%X \n","F")
	// str := "\u042F"
    // fmt.Println(str)

	// var k float64 = 123.456;
	// fmt.Printf("%f \n",k)
	// fmt.Printf("%e \n",k)

	
	port:= "5000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		text := "Hello world"
		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0:" + port

	log.Println("server starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}