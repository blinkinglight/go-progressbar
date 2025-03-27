package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func main() {
	r := chi.NewMux()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Main().Render(r.Context(), w)
	})

	r.Post("/progress", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		progress := 0
		total := 10124
		processed := 0
		for {
			select {
			case <-r.Context().Done():
				return
			default:
				processed++
				progress = (processed * 100 / total)
				time.Sleep(1 * time.Millisecond)
				sse.MergeFragmentTempl(Progress(fmt.Sprintf("%d", processed), fmt.Sprintf("%d", total), fmt.Sprintf("%d", progress)))
				if progress == 100 {
					sse.MergeFragmentTempl(Done())
					return
				}
			}
		}
	})

	log.Printf("Listening on :8089")
	http.ListenAndServe(":8089", r)
}
