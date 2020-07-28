package main

// Category type for storing podcast categories
type Category struct {
	Text           string     `xml:"text,attr"`
	ItunesCategory []Category `xml:"itunes:category"`
}
