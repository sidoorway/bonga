package data

type ChatPage struct {
	Lang    string
	Title   string
	Domain  string
	Color   string
	Genders []string
	Chat    bool
	Room    Room
	Rooms   []Room
}

type GenderSet struct {
	Title string
	Rooms []Room
}

type MainPage struct {
	Lang    string
	Title   string
	Domain  string
	Color   string
	Genders []GenderSet
}
