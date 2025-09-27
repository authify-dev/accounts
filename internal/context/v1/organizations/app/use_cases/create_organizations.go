package organizations_use_cases

import (
	"accounts/internal/context/v1/organizations/domain/commands"
	"accounts/internal/context/v1/organizations/domain/entities"
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

func (u *CreateOrganizationsUseCase) Execute(cc *customctx.CustomContext, command commands.CreateOrganizationCommand) utils.Response[entities.Organization] {

	res := u.repository.Save(command.ToEntity())

	if res.Err != nil {
		return utils.Response[entities.Organization]{
			Success:    false,
			StatusCode: res.Err.GetCode(),
			Error:      cc.NewError(res.Err),
		}
	}
	return utils.Response[entities.Organization]{
		Data:       res.Data,
		Success:    true,
		StatusCode: http.StatusCreated,
	}

}
