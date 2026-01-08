// https://api.keybit.ir/time/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "https://api.keybit.ir/time/"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic("create req issue")
	}

	res, er := client.Do(req)
	if er != nil {
		panic("request calling error")
	}
	defer res.Body.Close()
	body, e := io.ReadAll(res.Body)
	if e != nil {
		panic("read body issue")
	}

	fmt.Println(string(body))
}
