package download

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bernmarx/download-manager/internal/infrastructure/logger"
)

// Assembler combines multiple Packages
// into one complete file
type Assembler struct {
}

func NewAssembler() *Assembler {
	return &Assembler{}
}

func (a *Assembler) Combine(total int, pkg <-chan *Package) ([]byte, error) {
	dataParts := make([][]byte, total)

	for i := 0; i < total; i++ {
		p := <-pkg
		r := p.Response

		if (r.StatusCode != http.StatusOK) && (r.StatusCode != http.StatusPartialContent) {
			return nil, errors.New("got wrong code " + fmt.Sprint(r.StatusCode))
		}

		var err error
		dataParts[p.Idx], err = ioutil.ReadAll(r.Body)

		if err != nil {
			return nil, err
		}

		logger.Logger().Info("assembled " + fmt.Sprint(p.Idx) + " package")
	}

	data := make([]byte, 0)

	for i := 0; i < total; i++ {
		data = append(data, dataParts[i]...)
	}

	return data, nil
}
