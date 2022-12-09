package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs/v2"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translsateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", translsateUrl, nil)
	if err != nil {
		log.Fatalf("1. There was a problem: %s", err)
	}

	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("2. There was a problem: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, try again later"
		wg.Done()
		return
	}

	parsedJSON, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		log.Fatalf("3. There was a problem: %s", err)
	}

	nestOne, err := parsedJSON.ArrayElement(0)
	if err != nil {
		log.Fatalf("4. There was a problem: %s", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatalf("5. There was a problem: %s", err)
	}

	translatedStr, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatalf("6. There was a problem: %s", err)
	}

	str <- translatedStr.Data().(string)
	wg.Done()
}
