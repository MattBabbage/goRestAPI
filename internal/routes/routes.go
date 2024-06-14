package routes

import (
	"fmt"
	"net/http"

	"github.com/MattBabbage/goRestAPI/internal/weather"
)

func APIRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/getWeather", weatherHandler)
	mux.HandleFunc("/api/data", apiDataHandler)
	return mux
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, weather.GetWeather())
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from API"
	fmt.Fprint(w, data)
}
