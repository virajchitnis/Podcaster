package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Podcast type for holding the details of a podcast.
type Podcast struct {
	XMLName     xml.Name `xml:"channel"`
	ShortName   string   `yaml:"shortname"`
	Title       string   `xml:"title" yaml:"title"`
	ItunesTitle string   `xml:"itunes:title,omitempty" yaml:"itunes_title"`
	Description string   `xml:"description" yaml:"description"`
	Link        string   `xml:"link,omitempty" yaml:"link"`
	Image       struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url" yaml:"url"`
		Title   string   `xml:"title" yaml:"title"`
		Link    string   `xml:"link" yaml:"link"`
	} `yaml:"image"`
	ItunesImage      ItunesImage `yaml:"itunes_image"`
	Generator        string      `xml:"generator,omitempty"`
	LastBuildDate    string      `xml:"lastBuildDate"`
	Author           string      `xml:"author" yaml:"author"`
	ItunesAuthor     string      `xml:"itunes:author,omitempty" yaml:"itunes_author"`
	Copyright        string      `xml:"copyright,omitempty" yaml:"copyright"`
	Language         string      `xml:"language" yaml:"language"`
	ManagingEditor   string      `xml:"managingEditor,omitempty" yaml:"managing_editor"`
	WebMaster        string      `xml:"webMaster,omitempty" yaml:"webmaster"`
	Category         string      `xml:"category,omitempty" yaml:"category"`
	ItunesCategories []Category  `yaml:"itunes_categories"`
	ItunesSummary    string      `xml:"itunes:summary,omitempty" yaml:"itunes_summary"`
	ItunesType       ShowType    `xml:"itunes:type,omitempty" yaml:"itunes_type"`
	ItunesOwner      struct {
		XMLName     xml.Name `xml:"itunes:owner"`
		ItunesName  string   `xml:"itunes:name" yaml:"name"`
		ItunesEmail string   `xml:"itunes:email" yaml:"email"`
	} `yaml:"itunes_owner"`
	ItunesExplicit   YesNoType `xml:"itunes:explicit" yaml:"itunes_explicit"`
	Items            []Episode
	ItunesNewFeedURL string    `xml:"itunes:new-feed-url,omitempty" yaml:"new_feed_url"`
	ItunesBlock      YesNoType `xml:"itunes:block,omitempty" yaml:"itunes_block"`
	ItunesComplete   YesNoType `xml:"itunes:complete,omitempty" yaml:"itunes_complete"`
}

func (p *Podcast) addEpisode(episode Episode) {
	p.Items = append(p.Items, episode)
}

// PodcastReader type for reading podcast details from a yaml file
type PodcastReader struct {
	Podcast Podcast `yaml:"podcast"`
}

func readPodcastFromFile(filename string) Podcast {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file!")
		log.Fatal(err)
	}
	defer f.Close()

	var podcastReader PodcastReader
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&podcastReader)
	if err != nil {
		fmt.Println("Invalid file format!")
		log.Fatal(err)
	}
	podcastReader.Podcast.Generator = "https://github.com/virajchitnis/Podcaster"
	return podcastReader.Podcast
}
