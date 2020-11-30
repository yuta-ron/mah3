// Sample language-quickstart uses the Google Cloud Natural API to analyze the
// sentiment of "Hello, world!".
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	ctx := context.Background()
	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	flag.Parse()
	text := flag.Arg(0)

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	var dtl []string

	detected := false

	dtl = append(dtl, "<h1>Review警察👮‍♀️</h1>")

	dscore := sentiment.DocumentSentiment.GetScore()
	if dscore < 0.2 {
		detected = true

		fstr := fmt.Sprintf("文章全体のスコアが%fです。<br>文章全体が不穏かもです。カリカリしてたらクールダウンしましょう🐹", dscore)
		dtl = append(dtl, fstr)
	}

	dtl = append(dtl, "<table><thead><tr><th>問題の表現</th><th>感情分析スコア</th></tr></thead><tbody><tr>")

	var ss []string
	for _, v := range sentiment.GetSentences() {
		if v.Sentiment.GetScore() < 0 {
			fstr := fmt.Sprintf("<tr><td><code>%s</code></td><td>%f</td></tr>", v.GetText().GetContent(), v.Sentiment.GetScore())
			ss = append(ss, fstr)
		}
	}
	if len(ss) > 0 {
		detected = true
		dtl = append(dtl, "<h3>この言葉は受け手が傷付いたかもよ！</h3>")
		dtl = append(dtl, ss...)
	}

	dtl = append(dtl, "</tr></tbody></table>")

	if !detected {
		return
	}

	if detected {
		for _, v := range dtl {
			fmt.Print(v)
		}
	}
}

// //
// func tableHTML(st []string) string {

// 	return "a"
// }
