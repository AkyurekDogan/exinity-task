/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
)

// Rating interface provides partner rating methods
type Rating interface {
	Get(filter dto.Filter) (*dto.Rating, error)
}

type rating struct {
	dbPartnerRating repository.Rating
}

// NewRating creates a new partner rating service to access relevant operations
func NewRating(repoPartnerRating repository.Rating) Rating {
	return &rating{
		dbPartnerRating: repoPartnerRating,
	}
}

// Get returns the relavent partner rating data by filter
func (s *rating) Get(filter dto.Filter) (*dto.Rating, error) {
	eFilter := s.toModel(filter)
	partnerRating, err := s.dbPartnerRating.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDTO(*partnerRating)
	return &result, nil
}

func (s *rating) toDTO(pr model.Rating) dto.Rating {
	return dto.Rating{
		ValueAVG: pr.ValueAVG,
	}
}

func (s *rating) toModel(p dto.Filter) model.Filter {
	return model.Filter{
		PartnerId: p.PartnerId,
	}
}
