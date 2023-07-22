package recommendation

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/gorilla/mux"
)

type Handler struct {
	svc Service
}

type GetRecommendationResponse struct {
	HotelName string `json:"hotelName"`
	TotalCost struct {
		Cost     int64  `json:"cost"`
		Currency string `json:"currency"`
	} `json:"totalCost"`
}

func NewHandler(svc Service) (*Handler, error) {
	if svc == (Service{}) {
		return nil, errors.New("service cannot be empty")
	}

	return &Handler{svc: svc}, nil
}

func (h Handler) GetRecommendation(w http.ResponseWriter, req *http.Request) {
	// That supposed to work but it doesn't return any parameters for
	// curl -i --location --request GET 'http://localhost:4040/recommendation?location=UK&from=2022-09-01&to=2022-09-08&budget=5000'
	q := mux.Vars(req)

	location, ok := q["location"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	from, ok := q["from"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	to, ok := q["to"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	budget, ok := q["budget"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const expectedFormat = "2006-01-02"

	formattedStart, err := time.Parse(expectedFormat, from)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	formattedEnd, err := time.Parse(expectedFormat, to)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseInt(budget, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	budgetMon := money.New(b, "USD")

	rec, err := h.svc.Get(req.Context(), formattedStart, formattedEnd, location, budgetMon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(GetRecommendationResponse{
		HotelName: rec.HotelName,
		TotalCost: struct {
			Cost     int64  `json:"cost"`
			Currency string `json:"currency"`
		}{
			Cost:     rec.TripPrice.Amount(),
			Currency: "USD",
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
