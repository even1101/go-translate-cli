package cli

import (
	"log"
	"net/http"
	"sync"
)

const TranslateUrl = "https://translate.googleapis.com/translate_a/single"

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

func RequestTranslate(reqBody *RequestBody, strChan chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", TranslateUrl, nil)

	query := req.URL.Query()

	query.Add("client", "gtx")

	query.Add("sl", reqBody.SourceLang)
	query.Add("tl", reqBody.TargetLang)
	query.Add("dt", "t")
	query.Add("q", reqBody.SourceText)

	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatal("There was problem:%s", err)
	}

	client.Do(req)
}
