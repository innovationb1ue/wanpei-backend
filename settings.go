package main

type Settings struct {
	Secret string
	Addr   string
}

func NewSettings() *Settings {
	return &Settings{
		Secret: "ggbob",
		Addr:   "localhost: 8096",
	}
}
