/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
)

// PartnerSkill interface provides partner skills methods
type Skill interface {
	Get(filter dto.Filter) (*dto.Skill, error)
}

type skill struct {
	dbPartnerSkill repository.Skill
}

// NewSkill creates a new partner skills service to access relevant operations
func NewSkill(repoPartnerSkill repository.Skill) Skill {
	return &skill{
		dbPartnerSkill: repoPartnerSkill,
	}
}

// Get returns the relavent partner skill data by filter
func (s *skill) Get(filter dto.Filter) (*dto.Skill, error) {
	eFilter := s.toModel(filter)
	partnerSkill, err := s.dbPartnerSkill.Get(eFilter)
	if err != nil {
		return nil, err
	}
	result := s.toDTO(*partnerSkill)
	return &result, nil
}

func (s *skill) toDTO(pr model.Skill) dto.Skill {
	return pr.Skills
}

func (s *skill) toModel(p dto.Filter) model.Filter {
	return model.Filter{
		PartnerId: p.PartnerId,
	}
}
