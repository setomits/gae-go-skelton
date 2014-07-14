package myapp

import (
	"net/http"

	"myhandlers"
)

func init() {
	http.HandleFunc("/", myhandlers.RootPage)
}
