package util

import (
	"net/url"
)

func ValidateURL(validatinURL string) bool {
	_, err := url.ParseRequestURI(validatinURL)
	return err == nil
}
