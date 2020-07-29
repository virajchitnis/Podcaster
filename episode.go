package main

import "encoding/xml"

// Episode type for holding the details of podcast episodes.
type Episode struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	GUID        string   `xml:"guid"`
	Creator     string   `xml:"dc:creator"`
	Enclosure   struct {
		URL  string   `xml:"url,attr"`
		Size int      `xml:"length,attr"`
		Type FileType `xml:"type,attr"`
	} `xml:"enclosure"`
	Date              string      `xml:"pubDate"`
	ItunesSummary     string      `xml:"itunes:summary"`
	ItunesExplicit    YesNoType   `xml:"itunes:explicit"`
	ItunesDuration    int         `xml:"itunes:duration"`
	ItunesSeason      int         `xml:"itunes:season,omitempty"`
	ItunesEpisode     int         `xml:"itunes:episode,omitempty"`
	ItunesEpisodeType EpisodeType `xml:"itunes:episodeType"`
}

func (e *Episode) setDescription(desc string) {
	e.Description = desc
	e.ItunesSummary = desc
}
