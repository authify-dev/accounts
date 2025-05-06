package dtos

type GetTokenGoogleDTO struct {
	Code string `json:"code" binding:"required"`
}

func (dto GetTokenGoogleDTO) Validate() error {
	return nil
}
