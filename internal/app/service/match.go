/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"math"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
)

const (
	metricDistanceKM     = "km"
	metricCalcualtionAVG = "avg"
)

// Match interface provides methods for matching partner-customer
type Match interface {
	Find(filter dto.MatchFilter) (*dto.MatchListResponse, error)
}

type match struct {
	dbMatch repository.Match
}

// NewMatch creates a new match service for matching
func NewMatch(db repository.Match) Match {
	return &match{
		dbMatch: db,
	}
}

// Get returns the matching records by customer filter details
func (s *match) Find(filter dto.MatchFilter) (*dto.MatchListResponse, error) {
	flt := s.toModel(filter)
	data, err := s.dbMatch.Get(flt)
	if err != nil {
		return nil, err
	}
	domData := s.toDTOList(data)
	result := dto.MatchListResponse{
		Filter:  filter,
		Matches: domData,
	}
	return &result, nil
}

func (s *match) toModel(filter dto.MatchFilter) model.MatchFilter {
	return model.MatchFilter{
		MaterialType: filter.MaterialType,
		Loc: model.Location{
			Lat:  filter.Loc.Lat,
			Long: filter.Loc.Long,
		},
	}
}

func (s *match) toDTOList(eml model.MatchList) dto.MatchList {
	domML := make(dto.MatchList, 0, len(eml))
	for _, v := range eml {
		domML = append(domML, s.toDTO(v))
	}
	return domML
}

func (s *match) toDTO(m model.Match) dto.Match {
	return dto.Match{
		PartnerId: m.PartnerId,
		Name:      m.Name,
		Loc: dto.Location{
			Lat:  m.Loc.Lat,
			Long: m.Loc.Long,
		},
		Radius: dto.Measure{
			Value:  float32(m.Radius),
			Metric: metricDistanceKM,
		},
		Distance: dto.Measure{
			Value:  float32(math.Round((m.Distance/1000)*100) / 100), // convert to km and round to 2 digits
			Metric: metricDistanceKM,
		},
		Rating: dto.Rating{
			ValueAVG: m.Rating,
		},
		Skills: m.Skills,
		Rank:   m.Rank,
	}
}
