package urlshortener

import (
	"fmt"
	"math"
	"strconv"
)

type URLStore struct {
	urls map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string)
	}
}

func (s *URLStore) Get(key string) string {
	if url, ok := s.urls[key]; ok {
		return url
	}
	return ""
}

func (s *URLStore) Set(key, url string) bool {
	if _, ok := s.urls[key]; ok {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	key := genKey(s.Count())
	s.Set(key, url)
	return key
}

func genKey(n int) string {
	if n%2 == 0 {
		shortURL := "GoodKey№" + strconv.Itoa(n)
		return shortURL
	}	
	shortURL := "PoorKey№" + strconv.Itoa(n)
	return shortURL
}
