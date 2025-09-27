package organizations_use_cases

import (
	"accounts/internal/api/v1/organizations/domain/entities"
	organizations_gorm "accounts/internal/db/postgres/organinizations"
	"foundation/domain/customctx"
	"foundation/utils"
	"net/http"
)

type CreateOrganizationsUseCase struct {
	repository *organizations_gorm.OrganizationsPostgresRepository
}

func NewCreateOrganizationsUseCase(repository *organizations_gorm.OrganizationsPostgresRepository) *CreateOrganizationsUseCase {
	return &CreateOrganizationsUseCase{repository: repository}
}

func (u *CreateOrganizationsUseCase) Execute(cc *customctx.CustomContext) utils.Response[entities.Organization] {

	organization := entities.Organization{
		Name:       "Organization 1",
		RootUserID: "a1b21503-ca17-4188-9395-22f04da91f27",
	}

	res := u.repository.Save(organization)

	return utils.Response[entities.Organization]{
		Data:       res.Data,
		Success:    true,
		StatusCode: http.StatusCreated,
	}

}
