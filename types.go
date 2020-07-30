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

// FileType type for defining valid file types.
type FileType string

// Valid types for FileType
const (
	M4A       FileType = "audio/x-m4a"
	MPEG      FileType = "audio/mpeg"
	QuickTime FileType = "video/quicktime"
	MP4       FileType = "video/mp4"
	M4V       FileType = "video/x-m4v"
	PDF       FileType = "application/pdf"
)

// Category type for storing podcast categories
type Category struct {
	XMLName          xml.Name `xml:"itunes:category"`
	Text             string   `xml:"text,attr"`
	ItunesCategories []Category
}

func (c *Category) addCategory(cat Category) {
	c.ItunesCategories = append(c.ItunesCategories, cat)
}

// ItunesImage type for storing album artwork details
type ItunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    string   `xml:"href,attr" yaml:"href"`
}
