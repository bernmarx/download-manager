package download

import "net/http"

// Package represents a portion of
// a downloaded file
type Package struct {
	Response *http.Response
	Idx      int
}
