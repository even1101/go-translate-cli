package main

import (
	"flag"
	"fmt"
	"os"

	"go-translate-cli/cli"
)

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

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	cli.RequestTranslate(reqBody)
}
