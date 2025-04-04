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

type partnerTestCase struct {
	Name                   string
	RepoFilter             model.Filter
	RepoPartnerResponse    *model.Partner
	RepoPartnerError       error
	RepoSkillResponse      *model.Skill
	RepoSkillError         error
	RepoRatingResponse     *model.Rating
	RepoRatingError        error
	DTOFilter              dto.Filter
	DTOResponse            *dto.Partner
	HandlerSuccessResponse dto.Success
	HandlerErrorResponse   dto.Error
	Error                  error
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
			HandlerSuccessResponse: dto.Success{
				Response: dto.Response{
					StatusCode: 200,
					Message:    "OK",
				},
				Data: dto.Partner{
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
			},
			Error: nil,
		},
	}
}

func TestPartner_Get(t *testing.T) {
	prepPartnerSetup()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ratingRepositoryMock := repoMock.NewMockPartnerRating(ctrl)
	skillRepositoryMock := repoMock.NewMockPartnerSkill(ctrl)
	partnerRepositoryMock := repoMock.NewMockPartner(ctrl)

	ratingService := service.NewRating(ratingRepositoryMock)
	skillService := service.NewSkill(skillRepositoryMock)
	partnerService := service.NewPartner(partnerRepositoryMock, skillService, ratingService)

	// Create the handler
	partnerHandler := handler.NewPartner(partnerService)

	for _, v := range partnerTestCases {
		t.Run(v.Name, func(t *testing.T) {
			ratingRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoRatingResponse, v.RepoRatingError).Times(1)
			skillRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoSkillResponse, v.RepoSkillError).Times(1)
			partnerRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoPartnerResponse, v.RepoPartnerError).Times(1)

			// Create the request
			req := httptest.NewRequest("GET", fmt.Sprintf("/partner?id=%s", v.RepoFilter.PartnerId), nil)
			rec := httptest.NewRecorder()

			// Execute the handler
			partnerHandler.Get(rec, req)
			// Convert struct to JSON
			jsonBytes, err := json.Marshal(v.HandlerSuccessResponse)
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
