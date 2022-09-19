package server

type Settings struct {
	Secret     string
	Addr       string
	Sha256Salt string
	//Redis settings
	RedisAddr     string
	RedisPassword string

	// match making settings
	RedisMatchMakingUsersQueueName string
}

func NewSettings() *Settings {
	// define all the possible user settings here
	return &Settings{
		Secret:                         "ggbob",
		Addr:                           "localhost:8096",
		Sha256Salt:                     "salt123",
		RedisAddr:                      "localhost:6379",
		RedisPassword:                  "",
		RedisMatchMakingUsersQueueName: "match:users",
	}
}
