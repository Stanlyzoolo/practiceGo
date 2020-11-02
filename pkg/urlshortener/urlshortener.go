package urlshortener

import (
	"fmt"
	"math"
	"strconv"
)

type URLStore struct {
	urls map[string]string
}

type ErrNegativeSqrt float64

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
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

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Нельзя извлечь квадратный корень из отрицательного числа %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if math.IsNaN(math.Sqrt(x)) {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
