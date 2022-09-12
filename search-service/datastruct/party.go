package datastruct

import (
	"strings"
	"time"

	"github.com/clubo-app/clubben/protobuf/search"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Party struct {
	Id            string   `vespa:"-"`
	DocId         string   `vespa:"documentid,omitempty"`
	Title         string   `vespa:"title"`
	Description   string   `vespa:"description,omitempty"`
	MusicGenre    string   `vespa:"music_genre,omitempty"`
	Location      Location `vespa:"location"`
	EntryDate     int64    `vespa:"entry_date,omitempty"`
	IsPublic      bool     `vespa:"is_public"`
	FavoriteCount int32    `vespa:"favorite_count"`
}

type Location struct {
	Lat float32 `vespa:"lat"`
	Lng float32 `vespa:"lng"`
}

func (p Party) ToGRPCParty() *search.IndexedParty {
	pos := strings.LastIndex(p.DocId, ":")
	if pos == -1 {
		return &search.IndexedParty{}
	}

	return &search.IndexedParty{
		Id:            p.DocId[pos+1:],
		Title:         p.Title,
		Description:   p.Description,
		MusicGenre:    p.MusicGenre,
		Lat:           p.Location.Lat,
		Long:          p.Location.Lng,
		EntryDate:     timestamppb.New(time.Unix(p.EntryDate, 0)),
		IsPublic:      p.IsPublic,
		FavoriteCount: p.FavoriteCount,
	}
}
