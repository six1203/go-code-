package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	// json序列化的字段名
	Year int `json:"released"`
	// omitempty json如果为空可以省略该字段
	Color bool `json:"color,omitempty"`
	Actor []string
}

var movies = []Movie{
	{Title: "a", Year: 1942, Color: true, Actor: []string{"bob", "dylan", "tim"}},
	{Title: "b", Year: 1945, Color: false, Actor: []string{"dylan"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed:%v", err)
	}

	//fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
