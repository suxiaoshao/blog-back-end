package main

import (
	"blogServer/router"
)

func main() {
	r := router.Router()
	_ = r.Run("localhost:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
