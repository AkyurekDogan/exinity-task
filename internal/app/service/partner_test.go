/*
Service package handles the services for business logic and data processing
*/
package service_test

import (
	"testing"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	mock "github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/mock/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type partnerTestCase struct {
	Name                string
	RepoFilter          model.Filter
	RepoPartnerResponse *model.Partner
	RepoSkillResponse   *model.Skill
	RepoRatingResponse  *model.Rating
	RepoPartnerError    error
	RepoSkillError      error
	RepoRatingError     error
	DTOFilter           dto.Filter
	DTOResponse         *dto.Partner
	Error               error
}

var (
	partnerTestCases []partnerTestCase
)

func prepPartnerSetup() {
	partnerTestCases = []partnerTestCase{
		{
			Name: "Happy Path",
			RepoFilter: model.Filter{
				PartnerId: "1x1x1x1x",
			},
			RepoPartnerResponse: &model.Partner{
				Id:   "1x1x1x1",
				Name: "X Engineering",
				Loc: model.Location{
					Lat:  1,
					Long: 2,
				},
				Radius: 1,
			},
			RepoPartnerError: nil,
			RepoSkillResponse: &model.Skill{
				PartnerId: "1x1x1x1",
				Skills:    []string{"wood", "tile"},
			},
			RepoSkillError: nil,
			RepoRatingResponse: &model.Rating{
				PartnerId: "1x1x1x1",
				ValueAVG:  9,
			},
			RepoRatingError: nil,
			DTOFilter: dto.Filter{
				PartnerId: "1x1x1x1x",
			},
			DTOResponse: &dto.Partner{
				Id:   "1x1x1x1",
				Name: "X Engineering",
				Loc: dto.Location{
					Lat:  1,
					Long: 2,
				},
				Radius: dto.Measure{
					Value:  1,
					Metric: "km",
				},
				Rating: &dto.Rating{
					ValueAVG: 9,
				},
				Skills: &dto.Skill{"wood", "tile"},
			},
			Error: nil,
		},
	}
}

// TestPartnerGet ...
func TestPartnerGet(t *testing.T) {
	prepPartnerSetup()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ratingRepositoryMock := mock.NewMockPartnerRating(ctrl)
	skillRepositoryMock := mock.NewMockPartnerSkill(ctrl)
	partnerRepositoryMock := mock.NewMockPartner(ctrl)

	ratingService := service.NewRating(ratingRepositoryMock)
	skillService := service.NewSkill(skillRepositoryMock)

	partnerService := service.NewPartner(partnerRepositoryMock, skillService, ratingService)

	for _, v := range partnerTestCases {
		t.Run(v.Name, func(t *testing.T) {
			ratingRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoRatingResponse, v.RepoRatingError).Times(1)
			skillRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoSkillResponse, v.RepoSkillError).Times(1)
			partnerRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoPartnerResponse, v.RepoPartnerError).Times(1)

			partner, err := partnerService.Get(v.DTOFilter)
			// Assertions
			if v.Error == nil {
				assert.NoError(t, err)
				assert.Equal(t, v.DTOResponse, partner)
			} else {
				assert.Equal(t, v.Error.Error(), err.Error())
			}

		})
	}
}
