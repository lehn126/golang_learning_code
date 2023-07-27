package nettest

import (
	"fmt"
	"net/http"
	"work1/net/router"
)

func StartServices() {
	router.RegistHandlers()
	fmt.Println("http service start")
	e := http.ListenAndServe("localhost:8080", nil)
	if e != nil {
		fmt.Println("meet error when start http service:", e)
	}
}
