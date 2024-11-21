package conversation

import "github.com/go-playground/validator/v10"

type UpdateLastMessageDto struct {
	LastMessage string `gorm:"type:text" json:"last_message"`
}

func (c *UpdateLastMessageDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
