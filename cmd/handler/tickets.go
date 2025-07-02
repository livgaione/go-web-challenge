package handler

import (
	"log"
	"net/http"
	"tickets/internal/service"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type TicketHandler struct {
	service  service.ServiceTicket
	validate *validator.Validate
}

func NewTicketHandler(service service.ServiceTicket) *TicketHandler {
	v := validator.New()

	return &TicketHandler{
		service:  service,
		validate: v,
	}
}

func (h *TicketHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		s, err := h.service.GetAll(r.Context())

		if err != nil {

			response.JSON(w, http.StatusBadGateway, err.Error())

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": s,
		})

	}
}

func (h *TicketHandler) GetTicketByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		country := chi.URLParam(r, "country")
		log.Println("Country param:", country)

		s, err := h.service.GetTicketByDestinationCountry(country)

		if err != nil {

			response.JSON(w, http.StatusBadGateway, err.Error())

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": s,
		})

	}
}

func (h *TicketHandler) GetAverage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		country := chi.URLParam(r, "country")
		log.Println("Country param:", country)

		s, err := h.service.GetAverage(country)

		if err != nil {

			response.JSON(w, http.StatusBadGateway, err.Error())

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": s,
		})

	}
}
