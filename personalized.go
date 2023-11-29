package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

const RandomSeedCookieName = "randomSeed"
const RandomSeedLength = 32

func RequestSeed(w http.ResponseWriter, r *http.Request) int64 {
	seedCookie, err := r.Cookie(RandomSeedCookieName)
	if err == http.ErrNoCookie {
		seedCookie = &http.Cookie{
			Name:  RandomSeedCookieName,
			Value: strconv.FormatInt(rand.Int63(), 10),
		}
		http.SetCookie(w, seedCookie)
	}

	value, err := strconv.ParseInt(seedCookie.Value, 10, 64)
	if err != nil {
		return 0
	}

	return value
}

func ShuffleWindows(r *rand.Rand, windows []Window) {
	for i := range windows {
		j := r.Intn(i + 1)
		windows[i], windows[j] = windows[j], windows[i]
	}
}
