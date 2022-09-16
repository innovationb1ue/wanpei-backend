package server

type Settings struct {
	Secret     string
	Addr       string
	Sha256Salt string
}

func NewSettings() *Settings {
	return &Settings{
		Secret: "ggbob",
		Addr:   "localhost: 8096",
	}
}
