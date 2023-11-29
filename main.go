package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	loc, err := time.LoadLocation("Europe/Oslo")
	if err != nil {
		slog.Error("unable to load Oslo time zone", slog.Any("err", err))
	}

	staticFS := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFS))

	audioFS := http.FileServer(http.Dir("audio"))
	http.Handle("/audio/", http.StripPrefix("/audio/", blockOnDatePrefix(loc, audioFS)))

	http.HandleFunc("/", serveCalendar)

	slog.Info("listening for requests", slog.String("port", port))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func blockOnDatePrefix(loc *time.Location, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		day := time.Now().In(loc).Day()

		fileParts := strings.Split(r.URL.Path, "-")
		if len(fileParts) < 1 {
			return
		}

		fileDay, err := strconv.Atoi(fileParts[0])
		if err != nil {
			slog.Warn("unable to parse file day", slog.Any("err", err))
			return
		}

		if day < fileDay {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("forbidden! try again later :)"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}

type CalendarPageData struct {
	Title   string
	Windows []Window
}

type Window struct {
	Day       int
	Intro     string
	Song      string
	Placement string
}

func serveCalendar(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))

	windows := make([]Window, 24)
	for i := 0; i < 24; i++ {
		windows[i] = Window{
			Day:   i + 1,
			Intro: fmt.Sprintf("/audio/%02d-intro.mp3", i+1),
			Song:  fmt.Sprintf("/audio/%02d-song.mp3", i+1),
		}
	}

	seed := RequestSeed(w, r)
	personalRand := rand.New(rand.NewSource(seed))
	ShuffleWindows(personalRand, windows)

	pageData := CalendarPageData{
		Title:   "Hello World",
		Windows: windows,
	}

	err := t.Execute(w, pageData)
	if err != nil {
		slog.Error("unable to execute the template",
			slog.Any("err", err),
			slog.Time("time", time.Now()),
		)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
}
