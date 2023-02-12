package http

import (
	"net/http"
)

func getRawHttp(Uri string) *goquery.Document {
	log.Debug().Msg(fmt.Sprintf("Requested %s", Uri))
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(Uri)
	if err != nil {
		log.Fatal()
	}

	ReaderHTML, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal()
	}
	return ReaderHTML
}