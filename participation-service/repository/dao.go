package repository

import (
	"strings"

	"github.com/clubo-app/packages/cqlx"
	"github.com/go-playground/validator"
	"github.com/scylladb/gocqlx/v2"
)

type dao struct {
	sess *gocqlx.Session
}

func NewDB(keyspace, hosts string) (*gocqlx.Session, error) {
	h := strings.Split(hosts, ",")

	manager := cqlx.NewManager(keyspace, h)

	session, err := manager.Connect()
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func NewDAO(sess *gocqlx.Session) dao {
	return dao{sess: sess}
}

func (d *dao) NewInviteRepository(val *validator.Validate) InviteRepository {
	return &inviteRepository{sess: d.sess, val: val}
}

func (d *dao) NewParticipantRepository(val *validator.Validate) ParticipantRepository {
	return &participantRepository{sess: d.sess, val: val}
}
