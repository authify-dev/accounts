package organizations_use_cases

import (
	"accounts/internal/context/v1/organizations/domain/entities"
	organizations_gorm "accounts/internal/db/postgres/organizations"
	"accounts/internal/utils"
	"context"
)

type CreateOrganizationsUseCase struct {
	repository *organizations_gorm.OrganizationsPostgresRepository
}

func NewCreateOrganizationsUseCase(repository *organizations_gorm.OrganizationsPostgresRepository) *CreateOrganizationsUseCase {
	return &CreateOrganizationsUseCase{repository: repository}
}

func (u *CreateOrganizationsUseCase) Execute(ctx context.Context) utils.Responses[entities.Organization] {

	organization := entities.Organization{
		Name:       "Organization 1",
		RootUserID: "a1b21503-ca17-4188-9395-22f04da91f27",
	}

	res := u.repository.Save(organization)

	organization.ID = res.Data

	return utils.Responses[entities.Organization]{
		Body: organization,
	}

}
