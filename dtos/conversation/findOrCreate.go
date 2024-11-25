package conversation

import "github.com/go-playground/validator/v10"

type FindOrCreateDto struct {
	SenderID   uint `gorm:"not null" json:"sender_id"`
	ReceiverID uint `gorm:"not null" json:"receiver_id"`
}

func (c *FindOrCreateDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
