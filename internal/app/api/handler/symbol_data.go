/*
The Handler package to manage the request-response pipeline handlers
*/
package handler

import (
	"net/http"

	"github.com/AkyurekDogan/exinity-task/internal/app/service"
)

// SymbolData represents the SymbolData handler
type SymbolData interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type partner struct {
	base
	srvSymbolData service.SymbolData
}

// NewSymbolData returns the new partner service
func NewSymbolData(srvSymbolData service.SymbolData) SymbolData {
	return &partner{
		srvSymbolData: srvSymbolData,
	}
}

// @Summary The symbol data handler.
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
	/*
		 	err := s.srvSymbolData.Get()
			if err != nil {
				s.WriteErrorRespone(w, http.StatusInternalServerError, "internal server error", err)
				return
			}
			s.WriteSuccessRespone(w, http.StatusOK, nil)
	*/
}
