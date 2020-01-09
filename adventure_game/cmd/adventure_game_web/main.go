package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"problems_go/adventure_game"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the web app")
	filename := flag.String("file", "gopher.json", "JSON file with story")
	flag.Parse()
	fmt.Println(*filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := adventure_game.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := adventure_game.NewHandler(story)
	fmt.Printf("Start the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
