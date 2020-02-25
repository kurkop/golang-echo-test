package main

import (
	"github.com/kurkop/golang-echo-test/src/router"
)

// middlewares

func main() {
	e := router.New()
	e.Start("0.0.0.0:8000")
}
