package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://localhost:3000/videos", nil)
	req.Header.Add("Authorization", "Basic anVraGFuOnRlc3Q=")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Print(res)
	fmt.Println(string(body))
}
