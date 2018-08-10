package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/agg", enhancedAggHandler)

	http.ListenAndServe(":3000", nil)
}

func enhancedAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &s)

	news_map := make(map[string]NewsMap)

	var news_chan = make(chan News, len(s.Locations))

	for _, location := range s.Locations {
		wg.Add(1)
		go parse(location, news_chan)
	}

	wg.Wait()

	close(news_chan)

	for n := range news_chan {
		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing news aggregator", News: news_map}

	t, _ := template.ParseFiles(`basictemplate.html`)

	err := t.Execute(w, p)

	if err != nil {
		fmt.Println(err)
	}
}

func parse(location string, c chan<- News) {
	defer wg.Done()

	var n News

	siteResp, _ := http.Get(location)
	siteBytes, _ := ioutil.ReadAll(siteResp.Body)

	xml.Unmarshal(siteBytes, &n)

	c <- n
}
