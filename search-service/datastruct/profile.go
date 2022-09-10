package datastruct

import (
	"strings"

	"github.com/clubo-app/clubben/protobuf/search"
)

type Profile struct {
	Id          string `vespa:"-"`
	DocId       string `vespa:"documentid"`
	Username    string `vespa:"username"`
	Firstname   string `vespa:"firstname"`
	Lastname    string `vespa:"lastname"`
	FriendCount int32  `vespa:"friend_count"`
}

func (p Profile) ToGRPCProfile() *search.IndexedUser {
	pos := strings.LastIndex(p.DocId, ":")
	if pos == -1 {
		return &search.IndexedUser{}
	}

	return &search.IndexedUser{
		Id:        p.DocId[pos+1:],
		Username:  p.Username,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
	}
}
