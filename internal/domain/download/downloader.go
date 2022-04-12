package download

import (
	"fmt"
	"net/http"

	"github.com/bernmarx/download-manager/internal/infrastructure/logger"
)

type Downloader struct {
	httpgetter
}

type httpgetter interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewDownloader(hg httpgetter) *Downloader {
	return &Downloader{hg}
}

func (d *Downloader) Exec(req *http.Request, idx int, done chan<- *Package, errChan chan<- error) {
	logger.Logger().Info("executing download with idx " + fmt.Sprint(idx))
	resp, err := d.Do(req)

	if err != nil {
		errChan <- err
		return
	}

	done <- &Package{resp, idx}
}
