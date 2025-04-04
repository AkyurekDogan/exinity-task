/*
The Handler package to manage the request-response pipeline handlers
*/
package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AkyurekDogan/exinity-task/internal/app/dto"
	"github.com/AkyurekDogan/exinity-task/internal/app/handler"
	repoMock "github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/mock/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type matchTestCase struct {
	Name                   string
	RepoFilter             model.MatchFilter
	RepoResponse           model.MatchList
	ReporError             error
	DTOFilter              dto.MatchFilter
	DTOResponse            *dto.MatchListResponse
	HandlerResponseSuccess dto.Success
	HandlerResponseError   dto.Error
	Error                  error
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
			HandlerResponseSuccess: dto.Success{
				Response: dto.Response{
					StatusCode: 200,
					Message:    "OK",
				},
				Data: dto.MatchListResponse{
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
			},
		},
	}
}

func TestMatchGet(t *testing.T) {
	prepMatchSetup()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	matchRepositoryMock := repoMock.NewMockMatch(ctrl)
	matchService := service.NewMatch(matchRepositoryMock)
	// Create the handler
	partnerHandler := handler.NewMatch(matchService)

	for _, v := range matchTestCases {
		t.Run(v.Name, func(t *testing.T) {
			matchRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoResponse, v.ReporError).Times(1)
			// Create the request
			req := httptest.NewRequest("GET", fmt.Sprintf("/match?material_type=%s&lat=%f&long=%f", v.RepoFilter.MaterialType, v.RepoFilter.Loc.Lat, v.RepoFilter.Loc.Long), nil)
			rec := httptest.NewRecorder()

			// Execute the handler
			partnerHandler.Get(rec, req)
			// Convert struct to JSON
			jsonBytes, err := json.Marshal(v.HandlerResponseSuccess)
			if err != nil {
				fmt.Println("Error converting to JSON:", err)
				return
			}
			// Convert byte slice to string
			jsonString := string(jsonBytes)
			// Validate the response
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, jsonString, rec.Body.String())
		})
	}
}
