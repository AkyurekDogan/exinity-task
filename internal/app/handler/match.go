/*
The Handler package to manage the request-response pipeline handlers
*/
package handler

import (
	"net/http"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
)

// Match represents the match handler
type Match interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type match struct {
	base
	matchService service.Match
}

// NewMatch creates new match handler
func NewMatch(ms service.Match) Match {
	return &match{
		matchService: ms,
	}
}

// @Summary Returns the matched relevant partners for the customer
// @Description Regarding the customer parameters location and material type, the endpoint returns the matched partners with the
// rating descending and distance ascending orders
// @Tags match
// @Accept json
// @Produce json
// @Param material_type query string true "selected material type"
// @Param lat query string true "customer location as lattitute"
// @Param long query string true "customer location as longtitute"
// @Success 200 {object} dto.MatchListResponse "Success"
// @Failure 400 {object} dto.Error "Bad Request"
// @Failure 500 {object} dto.Error "Internal Server Error"
// @Router /match [get]
func (s *match) Get(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	filter, err := dto.NewMatchFilter(r)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid query parameters", err)
		return
	}
	// Validation check
	if err := filter.CheckFilter(); err != nil {
		s.WriteErrorRespone(w, http.StatusBadRequest, "invalid or insufficient input", err)
		return
	}
	result, err := s.matchService.Find(*filter)
	if err != nil {
		s.WriteErrorRespone(w, http.StatusInternalServerError, "internal server error", err)
		return
	}
	s.WriteSuccessRespone(w, http.StatusOK, result)
}
