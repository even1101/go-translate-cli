package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"go-translate-cli/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "zh-TW", "輸入的文字語系")
	flag.StringVar(&targetLang, "t", "en", "翻譯的文字語系")
	flag.StringVar(&sourceLang, "st", "Text", "翻譯的文字")

}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Print("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	strChan := make(chan string)

	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg)

	processedStr := strings.ReplaceAll(<-strChan, "+", " ")

	fmt.Printf("%s\n", processedStr)
	close(strChan)
	wg.Wait()
}
