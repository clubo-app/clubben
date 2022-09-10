package datastruct

type Party struct {
	Id       string `vespa:"-"`
	DocId    string `vespa:"documentid"`
	Title    string `vespa:"title"`
	IsPublic bool   `vespa:"isPublic"`
}
