/*
The dto package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package dto

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Filter filters the cost details
type Filter struct {
	PartnerId string `json:"id"`
}

// NewFilter initialize new filter from requested parameters
func NewFilter(r *http.Request) (*Filter, error) {
	var f Filter
	// get query parameters
	params := r.URL.Query()
	pPartnerId := params.Get("id")
	f.PartnerId = pPartnerId
	return &f, nil
}

// CheckFilter checks the filter field validations
func (f *Filter) CheckFilter() error {
	var valErrors []error
	if f.PartnerId == "" {
		valErrors = append(valErrors, errors.New("partner id must be provided"))
	}
	if len(valErrors) > 0 {
		return errors.Join(valErrors...)
	}
	return nil
}

// MatchFilter filters the cost details
type MatchFilter struct {
	MaterialType string   `json:"material_type"`
	Loc          Location `json:"location"`
}

// NewMatchFilter initialize new filter from requested parameters
func NewMatchFilter(r *http.Request) (*MatchFilter, error) {
	var f MatchFilter
	var errs []error
	// get query parameters
	params := r.URL.Query()
	pMaterialType := params.Get("material_type")
	pLat := params.Get("lat")
	pLong := params.Get("long")

	f.MaterialType = pMaterialType
	if strings.TrimSpace(pLat) != "" {
		num, err := strconv.ParseFloat(pLat, 64)
		if err != nil {
			errs = append(errs, errors.New("lattitute be provided as float64"))
		} else {
			f.Loc.Lat = num
		}
	}
	if strings.TrimSpace(pLong) != "" {
		num, err := strconv.ParseFloat(pLong, 64)
		if err != nil {
			errs = append(errs, errors.New("longtitute be provided as float64"))
		} else {
			f.Loc.Long = num
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return &f, nil
}

// CheckFilter checks the filter field validations
func (f *MatchFilter) CheckFilter() error {
	var valErrors []error
	if f.MaterialType == "" {
		valErrors = append(valErrors, errors.New("material type must be provided"))
	}
	if f.Loc.Lat <= 0 {
		valErrors = append(valErrors, errors.New("lattitute must be provided"))
	}
	if f.Loc.Long <= 0 {
		valErrors = append(valErrors, errors.New("longtitute must be provided"))
	}
	if len(valErrors) > 0 {
		return errors.Join(valErrors...)
	}
	return nil
}
