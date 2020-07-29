package main

import "encoding/xml"

// Podcast type for holding the details of a podcast.
type Podcast struct {
	XMLName     xml.Name `xml:"channel"`
	shortName   string
	Title       string `xml:"title"`
	ItunesTitle string `xml:"itunes:title,omitempty"`
	Description string `xml:"description"`
	Link        string `xml:"link,omitempty"`
	Image       struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	ItunesImage struct {
		XMLName xml.Name `xml:"itunes:image"`
		Href    string   `xml:"href,attr"`
	}
	Generator        string `xml:"generator,omitempty"`
	LastBuildDate    string `xml:"lastBuildDate"`
	Author           string `xml:"author"`
	ItunesAuthor     string `xml:"itunes:author,omitempty"`
	Copyright        string `xml:"copyright,omitempty"`
	Language         string `xml:"language"`
	ManagingEditor   string `xml:"managingEditor,omitempty"`
	WebMaster        string `xml:"webMaster,omitempty"`
	Category         string `xml:"category"`
	ItunesCategories []Category
	ItunesSummary    string   `xml:"itunes:summary"`
	ItunesType       ShowType `xml:"itunes:type,omitempty"`
	ItunesOwner      struct {
		XMLName     xml.Name `xml:"itunes:owner"`
		ItunesName  string   `xml:"itunes:name"`
		ItunesEmail string   `xml:"itunes:email"`
	}
	ItunesExplicit   YesNoType `xml:"itunes:explicit"`
	Items            []Episode
	ItunesNewFeedURL string    `xml:"itunes:new-feed-url,omitempty"`
	ItunesBlock      YesNoType `xml:"itunes:block,omitempty"`
	ItunesComplete   YesNoType `xml:"itunes:complete,omitempty"`
}

func (p *Podcast) setTitle(title string) {
	p.Title = title
	p.Image.Title = title
	p.ItunesTitle = title
}

func (p *Podcast) setDescription(desc string) {
	p.Description = desc
	p.ItunesSummary = desc
}

func (p *Podcast) setLink(link string) {
	p.Link = link
	p.Image.Link = link
}

func (p *Podcast) setImageURL(URL string) {
	p.Image.URL = URL
	p.ItunesImage.Href = URL
}

func (p *Podcast) setAuthor(author string) {
	p.Author = author
	p.ItunesAuthor = author
	p.ItunesOwner.ItunesName = author
}

func (p *Podcast) addCategory(cat Category) {
	p.ItunesCategories = append(p.ItunesCategories, cat)
}

func (p *Podcast) addEpisode(episode Episode) {
	p.Items = append(p.Items, episode)
}
