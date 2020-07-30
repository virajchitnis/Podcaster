package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
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
		r.GET("/"+podcast.shortName+"/feed.xml", func(c *gin.Context) {
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

func readConfigFile(config string) Config {
	f, err := os.Open(config)
	if err != nil {
		fmt.Println("Unable to read config file!")
		exitWithError()
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Invalid config file format!")
		exitWithError()
	}
	return cfg
}

func exitWithError() {
	os.Exit(1)
}

func buildPodcastList() {
	currTime := time.Now()

	newPodcast := Podcast{
		shortName:      "test",
		Copyright:      "© 2020 The Author. All Rights Reserved.",
		Language:       "en",
		Category:       "History",
		ItunesType:     Episodic,
		ItunesExplicit: No,
		LastBuildDate:  currTime.Format(time.RFC1123),
	}
	newPodcast.setTitle("TEST")
	newPodcast.setDescription("This is a test podcast")
	newPodcast.setAuthor("The Author & The Second Author")
	newPodcast.setImageURL("https://domain.com/image")
	newPodcast.setLink("https://twitter.com/blah")
	newPodcast.ItunesOwner.ItunesEmail = "blah@domain.com"
	newPodcast.addCategory(Category{
		Text: "History",
	})
	newsCategory := Category{
		Text: "News",
	}
	newsCategory.addCategory(Category{
		Text: "Politics",
	})
	newPodcast.addCategory(newsCategory)

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
