package news

import "log"

type Topic struct {
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

type newsAPIResponse struct {
	Articles []Topic `json:"articles"`
	Source   string  `json:"source"`
}

type Source struct {
	ID string `json:"id"`
}

type sourceAPIResponse struct {
	Status  string
	Sources []Source `json:"sources"`
}

type Archive map[string][]Topic

func NewArchive() Archive {
	return make(map[string][]Topic)
}

func (a Archive) data(category string) []Topic {
	return a[category]
}

func (a Archive) resetData() {
	for category := range a {
		delete(a, category)
	}
	log.Println("The archive is emptied.")
}

func (a Archive) collectNews(category string) {
	sources := getSourcesFor(category)
	topics := getLatest(sources)

	a[category] = topics
	log.Printf("The news for %s category are collected.\n", category)
}
