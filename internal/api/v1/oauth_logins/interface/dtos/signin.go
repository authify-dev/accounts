package dtos

type SigninGoogleDTO struct {
	Code string `json:"code" binding:"required"`
	Role string `json:"role" binding:"required"`
}

func (dto SigninGoogleDTO) Validate() error {
	return nil
}
