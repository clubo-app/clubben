package datastruct

type Party struct {
	Id            string   `vespa:"-"`
	DocId         string   `vespa:"documentid"`
	Title         string   `vespa:"title"`
	Description   string   `vespa:"description"`
	MusicGenre    string   `vespa:"music_genre"`
	Location      Location `vespa:"location"`
	EntryDate     int64    `vespa:"entry_date"`
	IsPublic      bool     `vespa:"is_public"`
	FavoriteCount int      `vespa:"favorite_count"`
}

type Location struct {
	Lat float32 `vespa:"lat"`
	Lng float32 `vespa:"lng"`
}
