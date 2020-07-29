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
		URL  string `xml:"url,attr"`
		Size int    `xml:"length,attr"`
		Type string `xml:"type,attr"`
	} `xml:"enclosure"`
	Date              string `xml:"pubDate"`
	ItunesSummary     string `xml:"itunes:summary"`
	ItunesExplicit    bool   `xml:"itunes:explicit"`
	ItunesDuration    string `xml:"itunes:duration"`
	ItunesSeason      int    `xml:"itunes:season"`
	ItunesEpisode     int    `xml:"itunes:episode"`
	ItunesEpisodeType string `xml:"itunes:episodeType"`
}
