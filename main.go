package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveCalendar)

	slog.Info("listening for requests", slog.String("port", port))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

type Window struct {
	Day       int
	Placement string
}

func serveCalendar(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))

	placements := []string{
		"top left",
		"top right",
		"bottom left",
		"bottom right",
	}

	windows := make([]Window, 24)
	for i := 0; i < 24; i++ {
		windows[i] = Window{
			Day:       i + 1,
			Placement: placements[rand.Intn(len(placements))],
		}
	}

	err := t.Execute(w, windows)
	if err != nil {
		slog.Error("unable to execute the template",
			slog.Any("err", err),
			slog.Time("time", time.Now()),
		)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
}
