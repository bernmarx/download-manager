package download

import (
	"net/http"
	"strconv"
)

// HeadSender sends HEAD request to
// the target URL and returns Bytes
// and Accept-Ranges
type HeadSender struct {
	HTTPSender
}

// Sender represents if a type can make
// HTTP requests
type HTTPSender interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHeadSender is a constructor for HeadSender
func NewHeadSender(sender HTTPSender) *HeadSender {
	return &HeadSender{sender}
}

// This a function for convinience, calling MakeRequest
// with default client
func MakeRequest(url string) (int64, bool, error) {
	h := NewHeadSender(http.DefaultClient)

	return h.MakeRequest(url)
}

// MakeRequest sends HEAD request and returns
// Content-Length and whether the target server
// accepts partial downloads
func (h *HeadSender) MakeRequest(url string) (int64, bool, error) {
	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return 0, false, err
	}

	resp, err := h.Do(req)

	if err != nil {
		return 0, false, err
	}

	cl := resp.Header.Get("Content-Length")

	//If cl is "" then server does not support
	//partial downloads
	if cl == "" {
		return 0, false, nil
	}

	cl64, err := strconv.ParseInt(cl, 0, 64)

	if err != nil {
		return 0, false, err
	}

	ar := resp.Header.Get("Accept-Ranges")

	if ar == "" {
		return cl64, false, nil
	}

	return cl64, true, nil
}
