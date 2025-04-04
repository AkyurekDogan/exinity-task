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

type skillTestCase struct {
	Name         string
	RepoFilter   model.Filter
	RepoResponse *model.Skill
	RepoError    error
	DTOFilter    dto.Filter
	DTOResponse  dto.Skill
	Error        error
}

var (
	skillTestCases []skillTestCase
)

func prepSkillSetup() {
	skillTestCases = []skillTestCase{
		{
			Name: "Happy Path",
			RepoFilter: model.Filter{
				PartnerId: "1x1x1x1x",
			},
			RepoResponse: &model.Skill{
				PartnerId: "1x1x1x1x",
				Skills:    []string{"wood", "tile"},
			},
			RepoError: nil,
			DTOFilter: dto.Filter{
				PartnerId: "1x1x1x1x",
			},
			DTOResponse: []string{"wood", "tile"},
			Error:       nil,
		},
		{
			Name: "Repository error",
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
func TestCostGet(t *testing.T) {
	prepSkillSetup()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	skillRepositoryMock := mock.NewMockPartnerSkill(ctrl)

	skillService := service.NewSkill(skillRepositoryMock)

	for _, v := range skillTestCases {
		t.Run(v.Name, func(t *testing.T) {
			skillRepositoryMock.EXPECT().Get(v.RepoFilter).Return(v.RepoResponse, v.RepoError).Times(1)
			skills, err := skillService.Get(v.DTOFilter)
			// Assertions
			if v.Error == nil {
				assert.NoError(t, err)
				assert.Equal(t, &v.DTOResponse, skills)
			} else {
				assert.Equal(t, v.Error.Error(), err.Error())
			}

		})
	}
}
