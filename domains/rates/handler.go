package rates

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/skobelina/currency_converter/utils/rest"
	"github.com/skobelina/currency_converter/utils/serializer"
)

func NewHandler(s Service) rest.Registrable {
	return &handler{s}
}

type handler struct {
	service Service
}

func (h *handler) Register(r *mux.Router) {
	r.HandleFunc("/api/rate", rest.ErrorHandler(h.get)).Methods("GET", "OPTIONS")
}

// swagger:route GET /rate Rate getRate
// Get the current USD to UAH rate
//
// responses:
//
//	200: body:rateResponse ok
//	400: statusBadRequest
func (h *handler) get(w http.ResponseWriter, r *http.Request) error {
	response, err := h.service.Get()
	if err != nil {
		return err
	}
	return serializer.SendJSON(w, http.StatusOK, response)
}
