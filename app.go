package main

import (
	f "fmt"
	h "net/http"
	"os"
	"time"
)

type Person struct {
	name, messege string
	age           int
}

func (p *Person) addPersonToFile() {
	date, _ := time.Now().UTC().MarshalText()
	os.WriteFile(p.name, []byte(p.messege+" "+string(date)), 0600)
	f.Printf("%s %s", "saved", p.name)
}

func handler(resp h.ResponseWriter, req *h.Request) {
	f.Fprintln(resp, "hello"+req.URL.Path[1:], req.URL.Path[1:])
}

func main() {
	p := Person{
		name:    "George",
		messege: "Hi George.",
		age:     21,
	}

	p.addPersonToFile()
	h.HandleFunc("/", handler)
	h.ListenAndServe(":8000", nil)
}
