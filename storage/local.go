package storage

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var quotes []string

func Prepare() error {
	file, err := os.Open("assets/quotes")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	return nil
}

func GetQuote() string {
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	return quotes[r.Intn(len(quotes))]
}
