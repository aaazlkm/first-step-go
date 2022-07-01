package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type GithubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil

}

func main() {
	// response, error := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	// if error != nil {
	// 	fmt.Println("Error:", error)
	// 	os.Exit(1)
	// }

	// customWriter := customWriter{}
	// io.Copy(customWriter, response.Body)

	// squares := square{10}
	// squares.doubleSize()
	// fmt.Println(squares.perimeter())

	// coloredS := coloredSquare{square{10}, "color"}
	// fmt.Println(coloredS.perimeter())

	// s := uppserString("aaaaa")
	// fmt.Println(s)
	// fmt.Println(s.upper())

	// ch := make(chan int)

}

func (s string) aa() {

}

type square struct {
	size int
}

type coloredSquare struct {
	square
	color string
}

func (s square) perimeter() (result int) {
	return s.size * 4
}

func (s *square) doubleSize() {
	s.size = s.size * 2
}

type uppserString string

func (u uppserString) upper() string {
	return strings.ToUpper(string(u))
}
