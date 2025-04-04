/*
Repository package for data access
*/
package repository

import (
	"database/sql"

	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/lib/pq"
)

// Skill represents the repository access layer for partner
type Skill interface {
	Get(f model.Filter) (*model.Skill, error)
}

type skill struct {
	driverRead *sql.DB
}

// NewSkill creates new database interface for PartnerSkill
func NewSkill(driverR *sql.DB) Skill {
	return &skill{
		driverRead: driverR,
	}
}

// Get gets the partner skill data
func (u *skill) Get(f model.Filter) (*model.Skill, error) {
	var result model.Skill
	// Execute a SELECT query
	err := u.driverRead.QueryRow(`
		select
			p.partner_id,
			p.craftsmanship_tags
		from public.skill p
		where partner_id=$1
	`, f.PartnerId).Scan(&result.PartnerId, pq.Array(&result.Skills))
	if err != nil {
		// Check for no rows found or other errors
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		} else {
			return nil, err
		}
	}
	return &result, nil
}
