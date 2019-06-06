package myshortener

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/kenshaw/baseconv"
	"log"
	"strconv"
	"strings"
)

var index int64 = 1

type MyShortener struct {
	storage map[string]string
}

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

func NewMyShortener() *MyShortener {
	return &MyShortener{storage: make(map[string]string)}
}

func hash(s string) string {

	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash

}

func getIndexInDigits62() string {
	index++
	s := strconv.FormatInt(index, 10)
	val62, ok := baseconv.Convert(s, baseconv.DigitsDec, baseconv.Digits62)
	if ok != nil {
		log.Printf("Cant covert index '%v' to Digits62\n", index)
		return ""
	}
	return val62

}

func split(u string) []string {
	var a []string
	a = strings.SplitN(u, "/", 2)
	if len(a) < 2 {
		log.Printf("func split arg u=" + string(u) + " - is not correct! \n Correct example: Otus.ru/contacts \n")
	}
	return a
}

func (s MyShortener) Shorten(url string) string {
	a := split(url)
	hashUrl := hash(url)
	if _, ok := s.storage[hashUrl]; ok {
		return split(s.storage[hashUrl])[1]
	} else {
		uri := getIndexInDigits62()
		shortUrl := a[0] + "/" + uri
		s.storage[hashUrl] = uri
		s.storage[uri] = url
		return shortUrl
	}

}

func (s MyShortener) Resolve(url string) string {
	a := split(url)
	u, ok := s.storage[a[1]]
	if !ok {
		return ""
	}
	return u
}
