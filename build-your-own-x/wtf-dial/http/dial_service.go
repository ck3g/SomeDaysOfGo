package http

import "net/url"

// DialService represents an HTTP implementation of wtf.DialService
type DialService struct {
	URL *url.URL
}
