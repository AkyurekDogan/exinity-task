/*
The Handler package to manage the request-response pipeline handlers
*/
package handler

import (
	"errors"
	"net/http"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
)

// Partner represents the partner handler
type Partner interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type partner struct {
	base
	partnerService service.Partner
}

// NewPartner returns the new partner service
func NewPartner(ps service.Partner) Partner {
	return &partner{
		partnerService: ps,
	}
}

// @Summary Returns the partner by partner_id parameter
// @Description Regarding the partner id parameter, returns the relevant partner with all details.
// @Tags partner
// @Accept json
// @Produce json
// @Param id query string true "patner_id"
// @Success 200 {object} dto.Partner "Success"
// @Failure 400 {object} dto.Error "Bad Request"
// @Failure 500 {object} dto.Error "Internal Server Error"
// @Router /partner [get]
func (s *partner) Get(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	filter, err := dto.NewFilter(r)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid query parameters", err)
		return
	}
	// Validation check
	if err := filter.CheckFilter(); err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid or insufficient input", err)
		return
	}
	result, err := s.partnerService.Get(*filter)
	if err != nil {
		if errors.Is(err, service.ErrNoPartner) {
			s.WriteErrorRespone(w, http.StatusBadRequest, "invalid parameters are provided", err)
			return
		}
		s.WriteErrorRespone(w, http.StatusInternalServerError, "internal server error", err)
		return
	}
	s.WriteSuccessRespone(w, http.StatusOK, result)
}
