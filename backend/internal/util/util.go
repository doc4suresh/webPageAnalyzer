package util

import (
	"net/url"
)

func ValidateURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	return err == nil
}
