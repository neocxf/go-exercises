package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"testing"
)

func TestAlternativeServer(t *testing.T) {
	runtime.GOMAXPROCS(10)

	http.HandleFunc("/agg", aggHandlerWithoutWg)

	http.ListenAndServe(":3000", nil)
}

func aggHandlerWithoutWg(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &s)

	newsMap := make(map[string]NewsMap)

	//numberOfWorkers := 10
	numberOfJobs := len(s.Locations)

	jobsChan := make(chan string)
	resultsChan := make(chan News)
	newsChan := make(chan News, numberOfJobs)
	doneSignal := make(chan struct{})

	for i := 0; i < numberOfJobs; i++ {
		go worker(jobsChan, resultsChan, doneSignal)
	}

	go func() {
		for _, location := range s.Locations {
			jobsChan <- location
		}
	}()

	go func() {
		count := 0
		for {
			newsChan <- <-resultsChan
			count++

			if count >= numberOfJobs {
				close(newsChan)
				doneSignal <- struct{}{}
				return
			}
		}
	}()

	<-doneSignal

	for n := range newsChan {

		for idx := range n.Keywords {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	resp2Client(w, newsMap)

}

func worker(jobs chan string, result chan News, done chan struct{}) {
	for {
		select {
		case loc := <-jobs:
			res, err := processOneLoc(loc)

			if err != nil {
				panic(err)
			}

			result <- res
		case <-done:
			return
		}
	}
}

func processOneLoc(location string) (News, error) {

	var n News

	siteResp, _ := http.Get(location)
	siteBytes, _ := ioutil.ReadAll(siteResp.Body)

	err := xml.Unmarshal(siteBytes, &n)

	return n, err
}

func resp2Client(w http.ResponseWriter, news_map map[string]NewsMap) {
	p := NewsAggPage{Title: "Amazing news aggregator", News: news_map}

	t, _ := template.ParseFiles(`basictemplate.html`)

	err := t.Execute(w, p)

	if err != nil {
		fmt.Println(err)
	}
}
