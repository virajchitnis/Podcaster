package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

var debug = flag.Bool("debug", false, "Enable debug mode")
var configFile = flag.String("config", "/etc/podcaster/config.yaml", "Configuration file")

var podcasts []Podcast

func main() {
	flag.Parse()
	config := readConfigFile(*configFile)

	fmt.Println("")
	fmt.Println("Welcome to Podcaster")
	fmt.Println("")

	if *debug != true {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Podcaster!",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Up!")
	})

	buildPodcastListFrom(config.Server.DataDirectory)

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
		r.GET("/"+podcast.ShortName+"/feed.xml", func(c *gin.Context) {
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

	r.Run(config.Server.Host + ":" + config.Server.Port)
}

func visitFile(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
				*files = append(*files, path)
			}
		}
		return nil
	}
}

func buildPodcastListFrom(directory string) {
	currTime := time.Now()

	var podcastFiles []string
	err := filepath.Walk(directory+"/podcasts", visitFile(&podcastFiles))
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range podcastFiles {
		newPodcast := readPodcastFromFile(file)

		newEpisode := Episode{
			GUID:              "hafuhgiahuha4r45",
			Creator:           "The Author & The Second Author",
			Date:              currTime.Format(time.RFC1123),
			ItunesExplicit:    No,
			ItunesEpisodeType: Full,
			ItunesDuration:    3256,
		}
		newEpisode.setTitle("Episode 1")
		newEpisode.setDescription("This is the first episode test.")
		newEpisode.Enclosure.URL = "https://www.domain.com/somefile.mp3"
		newEpisode.Enclosure.Size = 554543
		newEpisode.Enclosure.Type = MPEG
		newPodcast.addEpisode(newEpisode)
		podcasts = append(podcasts, newPodcast)
	}
}
