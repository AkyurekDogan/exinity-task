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

type matchTestCase struct {
	Name         string
	RepoFilter   model.MatchFilter
	RepoResponse model.MatchList
	ReporError   error
	DTOFilter    dto.MatchFilter
	DTOResponse  *dto.MatchListResponse
	Error        error
}

var (
	matchTestCases []matchTestCase
)

func prepMatchSetup() {
	matchTestCases = []matchTestCase{
		{
			Name: "Happy Path",
			RepoFilter: model.MatchFilter{
				MaterialType: "wood",
				Loc: model.Location{
					Lat:  1,
					Long: 1,
				},
			},
			RepoResponse: model.MatchList{
				{
					PartnerId: "1x1x1x1x",
					Name:      "X Engineering",
					Loc: model.Location{
						Lat:  1,
						Long: 1,
					},
					Radius:   1,
					Distance: 1000,
					Rating:   9,
					Skills:   []string{"wood", "tile"},
					Rank:     1,
				},
			},
			ReporError: nil,
			DTOFilter: dto.MatchFilter{
				MaterialType: "wood",
				Loc: dto.Location{
					Lat:  1,
					Long: 1,
				},
			},
			DTOResponse: &dto.MatchListResponse{
				Filter: dto.MatchFilter{
					MaterialType: "wood",
					Loc: dto.Location{
						Lat:  1,
						Long: 1,
					},
				},
				Matches: dto.MatchList{
					{
						PartnerId: "1x1x1x1x",
						Name:      "X Engineering",
						Loc: dto.Location{
							Lat:  1,
							Long: 1,
						},
						Radius: dto.Measure{
							Value:  1,
							Metric: "km",
						},
						Distance: dto.Measure{
							Value:  1,
							Metric: "km",
						},
						Rating: dto.Rating{
							ValueAVG: 9,
						},
						Skills: dto.Skill{"wood", "tile"},
						Rank:   1,
					},
				},
			},
			Error: nil,
		},
	}
}

// TestMatchGet ...
func TestMatchGet(t *testing.T) {
	prepMatchSetup()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	matchRepositoryMock := mock.NewMockMatch(ctrl)
	matchService := service.NewMatch(matchRepositoryMock)

	for _, v := range matchTestCases {
		t.Run(v.Name, func(t *testing.T) {
			matchRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoResponse, v.ReporError).Times(1)

			match, err := matchService.Find(v.DTOFilter)
			// Assertions
			if v.Error == nil {
				assert.NoError(t, err)
				assert.Equal(t, v.DTOResponse, match)
			} else {
				assert.Equal(t, v.Error.Error(), err.Error())
			}

		})
	}
}
