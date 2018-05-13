package xmltv

import (
	"time"
	"encoding/xml"
)

// Time that holds the time which is parsed from XML
type Time struct {
	time.Time
}

// UnmarshalXMLAttr used to marshal a time in the xmltv format to time in go format
func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	t1, err := time.Parse("20060102150405 -0700", attr.Value)
	if err != nil {
		return err
	}

	*t = Time{t1}
	return nil
}

// Tv the root element
type Tv struct {
	Channels          []Channel   `xml:"channel"                            json:"channels"`
	Programmes        []Programme `xml:"programme"                          json:"programmes"`
	Date              string      `xml:"date,attr,omitempty"                json:"date,omitempty"`
	SourceInfoURL     string      `xml:"source-info-url,attr,omitempty"     json:"source_info_url,omitempty"`
	SourceInfoName    string      `xml:"source-info-name,attr,omitempty"    json:"source_info_name,omitempty"`
	SourceDataURL     string      `xml:"source-data-url,attr,omitempty"     json:"source_data_url,omitempty"`
	GeneratorInfoName string      `xml:"generator-info-name,attr,omitempty" json:"generator_info_name,omitempty"`
	GeneratorInfoURL  string      `xml:"generator-info-url,attr,omitempty"  json:"generator_info_url,omitempty"`
}

// Channel details of a channel
type Channel struct {
	DisplayNames []CommonElement `xml:"display-name"   json:"display_names"`
	Icons        []Icon          `xml:"icon,omitempty" json:"icons,omitempty"`
	URLs         []string        `xml:"url,omitempty"  json:"urls,omitempty"`
	ID           string          `xml:"id,attr"        json:"id,omitempty"`
}

// Programme details of a single programme transmission
type Programme struct {
	Titles          []CommonElement  `xml:"title"                      json:"titles"`
	SecondaryTitles []CommonElement  `xml:"sub-title,omitempty"        json:"secondary_titles,omitempty"`
	Descriptions    []CommonElement  `xml:"desc,omitempty"             json:"descriptions,omitempty"`
	Credits         *Credits         `xml:"credits,omitempty"          json:"credits,omitempty"`
	Date            string           `xml:"date,omitempty"             json:"date,omitempty"`
	Categories      []CommonElement  `xml:"category,omitempty"         json:"categories,omitempty"`
	Keywords        []CommonElement  `xml:"keyword,omitempty"          json:"keywords,omitempty"`
	Languages       []CommonElement  `xml:"language,omitempty"         json:"languages,omitempty"`
	OrigLanguages   []CommonElement  `xml:"orig-language,omitempty"    json:"orig_languages,omitempty"`
	Length          *Length          `xml:"length,omitempty"           json:"length,omitempty"`
	Icons           []Icon           `xml:"icon,omitempty"             json:"icons,omitempty"`
	URLs            []string         `xml:"url,omitempty"              json:"urls,omitempty"`
	Countries       []CommonElement  `xml:"country,omitempty"          json:"countries,omitempty"`
	EpisodeNums     []EpisodeNum     `xml:"episode-num,omitempty"      json:"episode_nums,omitempty"`
	Video           *Video           `xml:"video,omitempty"            json:"video,omitempty"`
	Audio           *Audio           `xml:"audio,omitempty"            json:"audio,omitempty"`
	PreviouslyShown *PreviouslyShown `xml:"previously-shown,omitempty" json:"previously_shown,omitempty"`
	Premiere        *CommonElement   `xml:"premiere,omitempty"         json:"premiere,omitempty"`
	LastChance      *CommonElement   `xml:"last-chance,omitempty"      json:"last_chance,omitempty"`
	New             ElementPresent   `xml:"new"                        json:"new"`
	Subtitles       []Subtitle       `xml:"subtitles,omitempty"        json:"subtitles,omitempty"`
	Ratings         []Rating         `xml:"rating,omitempty"           json:"ratings,omitempty"`
	StarRatings     []Rating         `xml:"star-rating,omitempty"      json:"star_ratings,omitempty"`
	Reviews         []Review         `xml:"review,omitempty"           json:"reviews,omitempty"`
	Start           *Time            `xml:"start,attr"                 json:"start"`
	Stop            *Time            `xml:"stop,attr,omitempty"        json:"stop,omitempty"`
	PDCStart        *Time            `xml:"pdc-start,attr,omitempty"   json:"pdc_start,omitempty"`
	VPSStart        *Time            `xml:"vps-start,attr,omitempty"   json:"vps_start,omitempty"`
	Showview        string           `xml:"showview,attr,omitempty"    json:"showview,omitempty"`
	Videoplus       string           `xml:"videoplus,attr,omitempty"   json:"videoplus,omitempty"`
	Channel         string           `xml:"channel,attr"               json:"channel"`
	Clumpidx        string           `xml:"clumpidx,attr,omitempty"    json:"clumpidx,omitempty"`
}

// CommonElement element structure that is common, i.e. <country lang="en">Italy</country>
type CommonElement struct {
	Lang  string `xml:"lang,attr,omitempty" json:"lang,omitempty"`
	Value string `xml:",chardata"           json:"value,omitempty"`
}

// ElementPresent used to determine if element is present or not
type ElementPresent bool

// UnmarshalXML used to determine if the element is present or not. see https://stackoverflow.com/a/46516243
func (c *ElementPresent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	*c = true
	return nil
}

// Icon associated with the element that contains it
type Icon struct {
	Source string `xml:"src,attr"         json:"source"`
	Width  int    `xml:"width,omitempty"  json:"width,omitempty"`
	Height int    `xml:"height,omitempty" json:"height,omitempty"`
}

// Credits for the programme
type Credits struct {
	Directors    []string `xml:"director,omitempty"    json:"directors,omitempty"`
	Actors       []Actor  `xml:"actor,omitempty"       json:"actors,omitempty"`
	Writers      []string `xml:"writer,omitempty"      json:"writers,omitempty"`
	Adapters     []string `xml:"adapter,omitempty"     json:"adapters,omitempty"`
	Producers    []string `xml:"producer,omitempty"    json:"producers,omitempty"`
	Composers    []string `xml:"composer,omitempty"    json:"composers,omitempty"`
	Editors      []string `xml:"editor,omitempty"      json:"editors,omitempty"`
	Presenters   []string `xml:"presenter,omitempty"   json:"presenters,omitempty"`
	Commentators []string `xml:"commentator,omitempty" json:"commentators,omitempty"`
	Guests       []string `xml:"guest,omitempty"       json:"guests,omitempty"`
}

// Actor in a programme
type Actor struct {
	Role  string `xml:"role,attr,omitempty" json:"role,omitempty"`
	Value string `xml:",chardata"           json:"value"`
}

// Length of the programme
type Length struct {
	Units string `xml:"units,attr" json:"units"`
	Value string `xml:",chardata"  json:"value"`
}

// EpisodeNum of the programme
type EpisodeNum struct {
	System string `xml:"system,attr,omitempty" json:"system,omitempty"`
	Value  string `xml:",chardata"             json:"value"`
}

// Video details of the programme
type Video struct {
	Present string `xml:"present,omitempty" json:"present,omitempty"`
	Colour  string `xml:"colour,omitempty"  json:"colour,omitempty"`
	Aspect  string `xml:"aspect,omitempty"  json:"aspect,omitempty"`
	Quality string `xml:"quality,omitempty" json:"quality,omitempty"`
}

// Audio details of the programme
type Audio struct {
	Present string `xml:"present,omitempty" json:"present,omitempty"`
	Stereo  string `xml:"stereo,omitempty"  json:"stereo,omitempty"`
}

// PreviouslyShown When and where the programme was last shown, if known.
type PreviouslyShown struct {
	Start   string `xml:"start,attr,omitempty"   json:"start,omitempty"`
	Channel string `xml:"channel,attr,omitempty" json:"channel,omitempty"`
}

// Subtitle in a programme
type Subtitle struct {
	Language *CommonElement `xml:"language,omitempty"  json:"language,omitempty"`
	Type     string         `xml:"type,attr,omitempty" json:"type,omitempty"`
}

// Rating of a programme
type Rating struct {
	Value  string `xml:"value"                 json:"value"`
	Icons  []Icon `xml:"icon,omitempty"        json:"icons,omitempty"`
	System string `xml:"system,attr,omitempty" json:"system,omitempty"`
}

// Review of a programme
type Review struct {
	Value    string `xml:",chardata"          json:"value"`
	Type     string `xml:"type"               json:"type"`
	Source   string `xml:"source,omitempty"   json:"source,omitempty"`
	Reviewer string `xml:"reviewer,omitempty" json:"reviewer,omitempty"`
	Lang     string `xml:"lang,omitempty"     json:"lang,omitempty"`
}
