package service

type WebPageInfo struct {
	URL              string         `json:"url"`
	HtmlVersion      string         `json:"htmlVersion"`
	Title            string         `json:"title"`
	HeadCount        map[string]int `json:"headCount"`
	AccessbleLinks   []string       `json:"AccessbleLinks"`
	InaccessbleLinks []string       `json:"InaccessbleLinks"`
	IsLogingForm     bool           `json:"isLogingForm"`
}
