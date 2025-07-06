package service

type WebPageInfo struct {
	URL               string         `json:"url"`
	HTMLVersion       string         `json:"HTMLVersion"`
	Title             string         `json:"title"`
	HeadCount         map[string]int `json:"headCount"`
	AccessibleLinks   int            `json:"AccessibleLinks"`
	InAccessibleLinks int            `json:"InAccessibleLinks"`
	IsLoginForm       bool           `json:"IsLoginForm"`
}
