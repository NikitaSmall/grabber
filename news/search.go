package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getSourcesFor(category string) []Source {
	body := getRequest(sourceURL(category))

	var sourceAPI sourceAPIResponse
	json.Unmarshal(body, &sourceAPI)

	return sourceAPI.Sources
}

func getLatest(sources []Source) []Topic {
	var topics []Topic

	for _, source := range sources {
		body := getRequest(articleURL(source))

		var articleAPI newsAPIResponse
		json.Unmarshal(body, &articleAPI)

		topics = append(topics, articleAPI.Articles...)
	}

	return topics
}

func getRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic("Cannot get sources!")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic("Cannot get sources!")
	}
	resp.Body.Close()

	return body
}

func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", category)
}

func articleURL(source Source) string {
	return fmt.Sprintf("https://newsapi.org/v1/articles?source=%s&sortBy=latest&apiKey=bf9a4c68daf4453c9edd9fe70fbd3b6e", source.ID)
}
