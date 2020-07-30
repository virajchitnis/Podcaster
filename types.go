package main

import "encoding/xml"

// ShowType type for defining valid show types.
type ShowType string

// Valid types for ShowType
const (
	Episodic ShowType = "episodic"
	Serial   ShowType = "serial"
)

// EpisodeType type for defining valid episode types.
type EpisodeType string

// Valid types for EpisodeType
const (
	Full    EpisodeType = "full"
	Trailer EpisodeType = "trailer"
	Bonus   EpisodeType = "bonus"
)

// YesNoType type for defining valid types for things that require a yes or a no value.
type YesNoType string

// Valid types for YesNoType
const (
	Yes YesNoType = "Yes"
	No  YesNoType = "No"
)

// Category type for storing podcast categories
type Category struct {
	XMLName          xml.Name   `xml:"itunes:category"`
	Text             string     `xml:"text,attr" yaml:"text"`
	ItunesCategories []Category `yaml:"itunes_categories"`
}

// ItunesImage type for storing album artwork details
type ItunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    string   `xml:"href,attr" yaml:"href"`
}
