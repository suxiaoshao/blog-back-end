package main

import (
	"nextBlogServer/router"
)

func main() {
	r := router.Router()
	_ = r.Run("0.0.0.0:8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
