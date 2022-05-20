package nekoswrapper

type Endpoint struct {
	Format string `json:"format"`
	Min    string `json:"min"`
	Max    string `json:"max"`
}

type Resource struct {
	ArtistHref string `json:"artist_href"`
	ArtistName string `json:"artist_name"`
	SourceUrl  string `json:"source_url"`
	Url        string `json:"url"`
}

type Endpoints map[string]Endpoint

type Resources []Resource
