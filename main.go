package main

import (
	"fmt"
	"net/http"
)

func checkConnection() error {
	var err error
	cnt := 20

	for i := 0; i < cnt && (i > 0 && err != nil); i++ {
		_, err = http.Get("https://www.microsoft.com/")
	}

	return err
}

func main() {
	fmt.Println(checkConnection())
}
