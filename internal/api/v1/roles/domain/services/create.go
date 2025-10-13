package services

import (
	"accounts/internal/api/v1/roles/domain/commands"
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *RolesService) Create(cc *customctx.CustomContext, command commands.CreateRoleCommand) utils.Response[entities.Role] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Creating role")

	entity := command.ToEntity()

	res := s.repository.Save(entity)

	if res.Err != nil {
		entry.Error("Error creating role", "error", res.Err)
		return utils.Response[entities.Role]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error creating role: "+res.Err.Error(),
					"roles.create.error_creating_role",
				),
			),
			StatusCode: res.Err.GetCode(),
			Success:    false,
		}
	}

	return utils.Response[entities.Role]{
		Data:       res.Data,
		Success:    true,
		StatusCode: http.StatusCreated,
	}
}
