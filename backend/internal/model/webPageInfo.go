package model

type WebPageInfo struct {
	URL              string   `json:"url"`
	HtmlVersion      string   `json:"htmlVersion"`
	Title            string   `json:"title"`
	HeadCount        []int    `json:"headCount"`
	HeadLevels       []string `json:"headLevels"`
	AccessbleLinks   []string `json:"AccessbleLinks"`
	InaccessbleLinks []string `json:"InaccessbleLinks"`
	IsLogingForm     bool     `json:"isLogingForm"`
}
