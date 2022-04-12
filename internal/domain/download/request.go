package download

import (
	"fmt"
	"math"
	"net/http"
)

const bytesPerReq = 10000 // 10kb

type Request struct {
}

func PrepareReqs(url string, bytes int64, acceptRanges bool) ([]*http.Request, error) {
	r := &Request{}

	return r.PrepareReqs(url, bytes, acceptRanges)
}

func (r *Request) PrepareReqs(url string, bytes int64, acceptRanges bool) ([]*http.Request, error) {
	//If target server does not support partial downloads
	//we just return one request for the entire resource
	if bytes == 0 || !acceptRanges {
		reqs := make([]*http.Request, 1)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		reqs[0] = req

		return reqs, nil
	}

	reqsLen := int(math.Ceil(float64(bytes) / bytesPerReq))
	reqs := make([]*http.Request, reqsLen)

	for i, _ := range reqs {
		if i == len(reqs)-1 {
			tempReq, err := http.NewRequest("GET", url, nil)

			if err != nil {
				return nil, err
			}

			tempReq.Header.Set("Range", fmt.Sprintf("bytes=%v-", bytesPerReq*(int64(reqsLen)-1)))

			reqs[i] = tempReq

			break
		}

		tempReq, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		tempReq.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", bytesPerReq*int64(i), bytesPerReq*int64(i+1)-1))

		reqs[i] = tempReq
	}

	return reqs, nil
}
