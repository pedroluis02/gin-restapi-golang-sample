package main

import (
	"log"

	"github.com/pedroluis02/gin-restapi-golang-sample/src/api"
)

func main() {
	log.Println("GIN Restful API")

	api.NewAndRun()
}
