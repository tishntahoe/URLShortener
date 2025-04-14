package gateway

import (
	"context"
	"github.com/tishntahoe/UrlShortener/internal/storage"
	"net/http"
	"time"
)

func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//logger

		return
	}

}
func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//logger

		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key := r.PathValue("id")
	var st storage.StorageInterface
	link, err := st.Get(ctx, key)
	if err != nil {
		// logger
	}
	w.Write([]byte(link))
	return
}
