package task

import (
	"github.com/go-playground/validator/v10"
	"github.com/jonasala/delicious-friday/common"
)

type CreateUpdateRequest struct {
	WorkOrder   string `json:"work_order"`
	Description string `json:"description" validate:"required_without=WorkOrder"`
}

func (req *CreateUpdateRequest) Validate() []*common.ErrorResponse {
	validate := validator.New()
	var errors []*common.ErrorResponse

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &common.ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}

	return errors
}
