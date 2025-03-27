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

	r.Get("/progress", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		progress := 0
		for {
			select {
			case <-r.Context().Done():
				return
			default:
				progress++

				// heavy work simulation (100ms)
				time.Sleep(100 * time.Millisecond)

				sse.MergeFragmentTempl(Progress(fmt.Sprintf("%d", progress)))
				if progress == 100 {
					return
				}
			}
		}
	})

	log.Printf("Listening on :8089")
	http.ListenAndServe(":8089", r)
}
