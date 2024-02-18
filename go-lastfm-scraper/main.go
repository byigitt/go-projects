package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Print("Give me the LastFM username: ")
	var username string
	fmt.Scanln(&username)

	url := "https://last.fm/user/" + username

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	aboutMe := doc.Find(".about-me-sidebar p").Text()
	aboutMeText := ""

	if len(aboutMe) > 0 {
		aboutMeText = aboutMe
	} else {
		aboutMeText = "No information available"
	}

	fmt.Print("About me of the " + username + ": " + aboutMeText)
}