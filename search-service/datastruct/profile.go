package datastruct

type Profile struct {
	Id        string `vespa:"-"`
	Username  string `vespa:"username"`
	Firstname string `vespa:"firstname"`
	Lastname  string `vespa:"lastname"`
}
