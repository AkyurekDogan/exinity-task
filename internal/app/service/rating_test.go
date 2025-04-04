/*
Service package handles the services for business logic and data processing
*/
package service_test

import (
	"errors"
	"testing"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	mock "github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/mock/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type ratingTestCase struct {
	Name         string
	RepoFilter   model.Filter
	RepoResponse *model.Rating
	RepoError    error
	DTOFilter    dto.Filter
	DTOResponse  *dto.Rating
	Error        error
}

var (
	ratingTestCases []ratingTestCase
)

func prepRatingSetup() {
	ratingTestCases = []ratingTestCase{
		{
			Name: "Happy Path",
			RepoFilter: model.Filter{
				PartnerId: "1x1x1x1x",
			},
			RepoResponse: &model.Rating{
				PartnerId: "1x1x1x1x",
				ValueAVG:  9,
			},
			RepoError: nil,
			DTOFilter: dto.Filter{
				PartnerId: "1x1x1x1x",
			},
			DTOResponse: &dto.Rating{
				ValueAVG: 9,
			},
			Error: nil,
		},
		{
			Name: "Error Path",
			RepoFilter: model.Filter{
				PartnerId: "1x1x1x1x",
			},
			RepoResponse: nil,
			RepoError:    errors.New("repo-error"),
			DTOFilter: dto.Filter{
				PartnerId: "1x1x1x1x",
			},
			DTOResponse: nil,
			Error:       errors.New("repo-error"),
		},
	}
}

// TestCostGet ...
func TestRatingGet(t *testing.T) {
	prepRatingSetup()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ratingRepositoryMock := mock.NewMockPartnerRating(ctrl)

	ratingService := service.NewRating(ratingRepositoryMock)

	for _, v := range ratingTestCases {
		t.Run(v.Name, func(t *testing.T) {
			ratingRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoResponse, v.RepoError).Times(1)
			rating, err := ratingService.Get(v.DTOFilter)
			// Assertions
			if v.Error == nil {
				assert.NoError(t, err)
				assert.Equal(t, v.DTOResponse, rating)
			} else {
				assert.Equal(t, v.Error.Error(), err.Error())
			}

		})
	}
}
