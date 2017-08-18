package main

import (
	"fmt"
	"os"
	//"encoding/json"
	"net/http"
	"io/ioutil"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	url := os.Args[1]
	fmt.Print("Parsing url...\n- ", url, "\n")

	//m := Message{"Hello", "World", 300}
	//b, err := json.Marshal(m)

	//fmt.Print(b)
	//fmt.Print("\n")
	//fmt.Print(err)
	//fmt.Print("\n")

	//var ms Message

	//err2 := json.Unmarshal(b, &ms)

	//fmt.Print(err2, ms)

	resp, err := http.Get(url)

	print("Error: ", err, "\n")
	body, err := ioutil.ReadAll(resp.Body)
	print("Error: ", err, "\n")
	print(string(body), "\n")
}