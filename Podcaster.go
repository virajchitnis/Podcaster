package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

var podcasts []Podcast

func main() {
	fmt.Println("Welcome to Podcaster!")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Podcaster!",
		})
	})

	buildPodcastList()

	for _, podcast := range podcasts {
		type XMLRoot struct {
			XMLName     xml.Name `xml:"rss"`
			ItunesNS    string   `xml:"xmlns:itunes,attr"`
			DCNS        string   `xml:"xmlns:dc,attr"`
			ContentNS   string   `xml:"xmlns:content,attr"`
			AtomNS      string   `xml:"xmlns:atom,attr"`
			Version     string   `xml:"version,attr"`
			PodcastData Podcast
		}
		r.GET("/"+podcast.shortName, func(c *gin.Context) {
			c.XML(http.StatusOK, XMLRoot{
				ItunesNS:    "http://www.itunes.com/dtds/podcast-1.0.dtd",
				DCNS:        "http://purl.org/dc/elements/1.1/",
				ContentNS:   "http://purl.org/rss/1.0/modules/content/",
				AtomNS:      "http://www.w3.org/2005/Atom",
				Version:     "2.0",
				PodcastData: podcast,
			})
		})
	}

	r.Run()
}

func buildPodcastList() {
	newPodcast := Podcast{
		shortName:   "test",
		Title:       "TEST",
		Description: "This is a test podcast.",
	}
	newPodcast.setImageURL("blah")
	podcasts = append(podcasts, newPodcast)
}
