package main

import (
	"github.com/kihoonkim/gogo/webserver"
)

func main() {
	webserver.InitRouter()
	webserver.Run()
}