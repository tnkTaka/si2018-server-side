package repositories

import (
	"log"

	"github.com/eure/si2018-server-side/entities"
	"github.com/go-openapi/strfmt"
	"github.com/go-xorm/builder"
)

type UserMessageRepository struct{}

func NewUserMessageRepository() UserMessageRepository {
	return UserMessageRepository{}
}

func (r *UserMessageRepository) Create(ent entities.UserMessage) error {
	s := engine.NewSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserMessageRepository) GetMessages(userID, partnerID int64, limit int, latest, oldest *strfmt.DateTime) ([]entities.UserMessage, error) {
	var messages []entities.UserMessage
	var ids = []int64{userID, partnerID}

	s := engine.NewSession()
	defer func() { log.Println(s.LastSQL()) }()
	s.Where(builder.In("user_id", ids))
	s.And(builder.In("partner_id", ids))
	if latest != nil {
		s.And("created_at > ?", oldest)
	}
	if oldest != nil {
		s.And("created_at < ?", latest)
	}
	s.Limit(limit)
	err := s.Find(&messages)
	if err != nil {
		return messages, err
	}

	return messages, nil
}