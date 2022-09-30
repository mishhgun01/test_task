package pkg

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// структура API
type API struct {
	r     *mux.Router
	cache *Storage
}

func New(r *mux.Router, s *Storage) *API {
	return &API{r: r, cache: s}
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/gas", api.gasStats).Methods(http.MethodGet, http.MethodPost)
}

func (api *API) ListenAndServe(addr string) error {
	err := http.ListenAndServe(addr, api.r)
	return err
}

func (api *API) gasStats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// получение статистики и отправка данных в кэш
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var output GasStatistics
		for _, transaction := range req.Data {
			for _, items := range transaction {
				output, err = statistics(items)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
		api.cache.UpdateData(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodGet:
		// отправка статистики, которая берется из кэша
		data, err := api.cache.GetData()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(output)
	}
}
