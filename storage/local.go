package storage

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var quotes []string
var users map[int64]*dummyUser // Will be replaced by redis

type dummyUser struct { // This one too
	IsAskingWeather bool
}

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

	users = make(map[int64]*dummyUser)

	return nil
}

func GetQuote() string {
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	return quotes[r.Intn(len(quotes))]
}

func GetUser(id int64) *dummyUser {
	if user, ok := users[id]; ok {
		return user
	}

	user := dummyUser{}
	users[id] = &user
	return &user
}
