/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"errors"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
)

var (
	ErrNoPartner = errors.New("partner is not found with given parameters")
)

// Partner interface provides price methods
type Partner interface {
	Get(filter dto.Filter) (*dto.Partner, error)
}

type partner struct {
	dbPartner repository.Partner
	svcSkill  Skill
	svcRating Rating
}

// NewPartner creates a new price service with following operations: get
func NewPartner(repoPartner repository.Partner,
	svcSkill Skill,
	svcRating Rating,
) Partner {
	return &partner{
		dbPartner: repoPartner,
		svcSkill:  svcSkill,
		svcRating: svcRating,
	}
}

// Get returns the relavent partner data
func (s *partner) Get(filter dto.Filter) (*dto.Partner, error) {
	eFilter := s.toModel(filter)
	partner, err := s.dbPartner.Get(eFilter)
	if err != nil {
		if errors.Is(err, repository.ErrNoRows) {
			return nil, ErrNoPartner
		}
		return nil, err
	}
	result := s.toDTO(partner)
	err = s.getPartnerSkills(&result, filter)
	if err != nil {
		return nil, err
	}
	err = s.getPartnerRating(&result, filter)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *partner) getPartnerSkills(p *dto.Partner, f dto.Filter) error {
	skill, err := s.svcSkill.Get(f)
	if err != nil {
		if !errors.Is(err, repository.ErrNoRows) {
			return err
		}
	}
	p.Skills = skill

	return nil
}

func (s *partner) getPartnerRating(p *dto.Partner, f dto.Filter) error {
	rating, err := s.svcRating.Get(f)
	if err != nil {
		if !errors.Is(err, repository.ErrNoRows) {
			return err
		}
	}
	p.Rating = rating
	return nil
}
func (s *partner) toDTO(p *model.Partner) dto.Partner {
	return dto.Partner{
		Id:   p.Id,
		Name: p.Name,
		Loc: dto.Location{
			Lat:  p.Loc.Lat,
			Long: p.Loc.Long,
		},
		Radius: dto.Measure{
			Value:  float32(p.Radius),
			Metric: metricDistanceKM,
		},
	}
}

func (s *partner) toModel(p dto.Filter) model.Filter {
	return model.Filter{
		PartnerId: p.PartnerId,
	}
}
