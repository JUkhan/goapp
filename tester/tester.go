package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JUkhan/goapp/util"
)

func xmain() {
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
func main() {
	sum := util.New([]util.T{1, 2, 3}).Map(func(item util.T, i int) util.T {
		return item.(int) * 2
	}).Filter(func(item util.T, i int) bool {
		return item.(int) < 6
	}).Foreach(func(t util.T) {
		fmt.Println(t)
	}).Map(func(t util.T, i int) util.T {
		return t.(int) * 10
	}).Foreach(func(t util.T) {
		fmt.Println(t)
	}).Scan(0, func(t1, t2 util.T, i int) util.T {
		return t1.(int) + t2.(int)
	}).Foreach(func(t util.T) {
		fmt.Println(t)
	}).Reduce(0, func(t1, t2 util.T, i int) util.T {
		return t1.(int) + t2.(int)
	})
	fmt.Println(sum)

}
