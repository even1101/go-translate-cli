package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

const TranslateUrl = "https://translate.googleapis.com/translate_a/single"

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

func RequestTranslate(reqBody *RequestBody, str chan string, wg *sync.WaitGroup) {
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
		log.Fatal("1. There was problem:%s ", err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("2. There was problem:%s ", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, try again later."
		wg.Done()
		return
	}

	parserJson, err := gabs.ParseJSONBuffer(res.Body)

	if err != nil {
		log.Fatal("3. There was problem:%s ", err)
	}

	nestOne, err := parserJson.ArrayElement(0)

	if err != nil {
		log.Fatal("4. There was problem:%s ", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)

	if err != nil {
		log.Fatal("5. There was problem:%s ", err)
	}

	translateStr, err := nestTwo.ArrayElement(0)

	if err != nil {
		log.Fatal("6. There was problem:%s ", err)
	}

	str <- translateStr.Data().(string)
	wg.Done()
}
