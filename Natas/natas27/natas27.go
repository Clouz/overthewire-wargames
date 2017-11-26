package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	result("?x1=1&y1=1&x2=300&y2=300")
}

func result(xx string) bool {

	usr := "natas27"
	psw := "55TBjpPZUUJgVP5b3BnbG6ON9uDPVzCJ"
	baseURL := "http://natas27.natas.labs.overthewire.org/"

	query := xx

	client := &http.Client{}
	req, _ := http.NewRequest("GET",
		baseURL+query, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(usr, psw))

	req.Header.Set("User-Agent", "Clouz>")
	cookie := http.Cookie{Name: "PHPSESSID", Value: "Clouz"}
	req.AddCookie(&cookie)

	res, _ := client.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bs := buf.String()
	fmt.Println("\n", res.Header, "\n")

	fmt.Println(bs)

	if strings.Contains(bs, "The credentials for the next level are") {
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
