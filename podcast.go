package main

// Podcast type for holding the details of a podcast.
type Podcast struct {
	shortName   string
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Image       struct {
		URL   string `xml:"url"`
		Title string `xml:"title"`
		Link  string `xml:"link"`
	} `xml:"image"`
	ItunesImage    string     `xml:"itunes:image"`
	Generator      string     `xml:"generator"`
	LastBuildDate  string     `xml:"lastBuildDate"`
	Author         string     `xml:"author"`
	ItunesAuthor   string     `xml:"itunes:author"`
	Copyright      string     `xml:"copyright"`
	Language       string     `xml:"language"`
	ManagingEditor string     `xml:"managingEditor"`
	WebMaster      string     `xml:"webMaster"`
	Category       string     `xml:"category"`
	ItunesCategory []Category `xml:"itunes:category"`
	ItunesSummary  string     `xml:"itunes:summary"`
	ItunesType     string     `xml:"itunes:type"`
	ItunesOwner    struct {
		ItunesName  string `xml:"itunes:name"`
		ItunesEmail string `xml:"itunes:email"`
	} `xml:"itunes:owner"`
	ItunesExplicit bool      `xml:"itunes:explicit"`
	Items          []Episode `xml:"item"`
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
	p.ItunesImage = URL
}

func (p *Podcast) setAuthor(author string) {
	p.Author = author
	p.ItunesAuthor = author
	p.ItunesOwner.ItunesName = author
}
