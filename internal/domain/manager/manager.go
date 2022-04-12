package manager

import (
	"net/http"
	"path/filepath"

	"github.com/bernmarx/download-manager/internal/domain/destination"
	"github.com/bernmarx/download-manager/internal/domain/download"
)

// Just one big function to make full download and write a file
func Download(url string, dest string) error {
	bytes, ar, err := download.MakeRequest(url)

	if err != nil {
		return err
	}

	reqs, err := download.PrepareReqs(url, bytes, ar)

	if err != nil {
		return err
	}

	errChan := make(chan error)
	done := make(chan *download.Package, len(reqs))

	for i := 0; i < len(reqs); i++ {
		d := download.NewDownloader(http.DefaultClient)
		go d.Exec(reqs[i], i, done, errChan)
	}

	a := download.NewAssembler()
	data, err := a.Combine(len(reqs), done)

	if err != nil {
		return err
	}

	d, err := destination.NewDestination(dest)

	if err != nil {
		return err
	}

	filename := filepath.Base(url)

	return d.WriteFile(filename, data)
}
