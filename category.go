package main

import "encoding/xml"

// Category type for storing podcast categories
type Category struct {
	XMLName        xml.Name `xml:"itunes:category"`
	Text           string   `xml:"text,attr"`
	ItunesCategory []Category
}
