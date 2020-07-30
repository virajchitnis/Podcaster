package main

import "encoding/xml"

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

// PodcastReader type for reading podcast details from a yaml file
type PodcastReader struct {
	Podcast Podcast `yaml:"podcast"`
}
