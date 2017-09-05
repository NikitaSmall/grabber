package main

import (
	"grabber/news"
	"grabber/router"
)

func main() {
	a := news.NewArchive()
	r := router.New()

	go a.CollectNews()
	r.Run()
}
