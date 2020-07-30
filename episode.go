package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Episode type for holding the details of podcast episodes.
type Episode struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title" yaml:"title"`
	ItunesTitle string   `xml:"itunes:title,omitempty" yaml:"itunes_title"`
	Description string   `xml:"description,omitempty" yaml:"description"`
	GUID        string   `xml:"guid,omitempty" yaml:"guid"`
	Creator     string   `xml:"dc:creator,omitempty" yaml:"creator"`
	Enclosure   struct {
		XMLName xml.Name `xml:"enclosure"`
		URL     string   `xml:"url,attr"`
		Size    int64    `xml:"length,attr"`
		Type    string   `xml:"type,attr"`
		File    string   `xml:"-" yaml:"file"`
	} `yaml:"enclosure"`
	Date              string        `xml:"pubDate,omitempty" yaml:"release_date"`
	ItunesSummary     string        `xml:"itunes:summary,omitempty" yaml:"itunes_summary"`
	ItunesExplicit    YesNoType     `xml:"itunes:explicit,omitempty" yaml:"itunes_explicit"`
	ItunesDuration    int           `xml:"itunes:duration,omitempty"`
	ItunesSeason      int           `xml:"itunes:season,omitempty" yaml:"season"`
	ItunesEpisode     int           `xml:"itunes:episode,omitempty" yaml:"episode"`
	ItunesEpisodeType EpisodeType   `xml:"itunes:episodeType,omitempty" yaml:"episode_type"`
	Link              string        `xml:"link,omitempty" yaml:"link"`
	ItunesImage       []ItunesImage `yaml:"itunes_image"`
	ItunesBlock       YesNoType     `xml:"itunes:block,omitempty" yaml:"itunes_block"`
}

func (e *Episode) readMediaFileDetails(basePath string) {
	mediaFileInfo, err := os.Lstat(basePath + "/" + e.Enclosure.File)
	if err != nil {
		log.Fatal(err)
	}
	e.Enclosure.Size = mediaFileInfo.Size()

	f, err := os.Open(basePath + "/" + e.Enclosure.File)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	e.Enclosure.Type = getFileContentType(f)
}

func (e *Episode) validatePubDate() {
	_, err := time.Parse(time.RFC1123, e.Date)
	if err != nil {
		log.Fatal(err)
	}
}

func getFileContentType(file *os.File) string {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	contentType := http.DetectContentType(buffer)

	return contentType
}

// EpisodeReader type for reading podcast details from a yaml file
type EpisodeReader struct {
	Episode Episode `yaml:"episode"`
}

func readEpisodeFromFile(filename string) Episode {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file!")
		log.Fatal(err)
	}
	defer f.Close()

	var episodeReader EpisodeReader
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&episodeReader)
	if err != nil {
		fmt.Println("Invalid file format!")
		log.Fatal(err)
	}
	return episodeReader.Episode
}
