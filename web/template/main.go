package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)

	http.ListenAndServe(":3000", nil)

}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex


	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &s)

	news_map := make(map[string]NewsMap)

	for _, location := range s.Locations {
		wg.Add(1)
		go parseLocation(location, news_map)
	}

	wg.Wait()

	p := NewsAggPage{Title:"Amazing news aggregator", News:news_map}

	t, _ := template.ParseFiles(`basictemplate.html`)

	err := t.Execute(w, p)

	if err != nil {
		fmt.Println(err)
	}
}


func parseLocation(location string, news_map map[string]NewsMap) {
	defer wg.Done()

	var n News

	siteResp, _ :=  http.Get(location)
	siteBytes,_ := ioutil.ReadAll(siteResp.Body)

	xml.Unmarshal(siteBytes, &n)

	for idx, _ := range n.Keywords {
		news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
	}


}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, welcome to go template</h1>")
}
