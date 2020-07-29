package main

import "encoding/xml"

// Episode type for holding the details of podcast episodes.
type Episode struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	ItunesTitle string   `xml:"itunes:title,omitempty"`
	Description string   `xml:"description,omitempty"`
	GUID        string   `xml:"guid,omitempty"`
	Creator     string   `xml:"dc:creator"`
	Enclosure   struct {
		XMLName xml.Name `xml:"enclosure"`
		URL     string   `xml:"url,attr"`
		Size    int      `xml:"length,attr"`
		Type    FileType `xml:"type,attr"`
	}
	Date              string      `xml:"pubDate,omitempty"`
	ItunesSummary     string      `xml:"itunes:summary"`
	ItunesExplicit    YesNoType   `xml:"itunes:explicit,omitempty"`
	ItunesDuration    int         `xml:"itunes:duration,omitempty"`
	ItunesSeason      int         `xml:"itunes:season,omitempty"`
	ItunesEpisode     int         `xml:"itunes:episode,omitempty"`
	ItunesEpisodeType EpisodeType `xml:"itunes:episodeType,omitempty"`
	Link              string      `xml:"link,omitempty"`
	ItunesImage       []ItunesImage
	ItunesBlock       YesNoType `xml:"itunes:block,omitempty"`
}

func (e *Episode) setTitle(title string) {
	e.Title = title
	e.ItunesTitle = title
}

func (e *Episode) setDescription(desc string) {
	e.Description = desc
	e.ItunesSummary = desc
}

func (e *Episode) addImage(img ItunesImage) {
	var blankSlice []ItunesImage
	e.ItunesImage = append(blankSlice, img)
}
