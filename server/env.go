package server

import (
	"os"
	"strings"
)

type Env struct {
	Vars map[string]string
}

func GetEnv() *Env {
	envVars := make(map[string]string)
	for _, envStr := range os.Environ() {
		pair := strings.SplitN(envStr, "=", 2)
		envVars[pair[0]] = pair[1]
	}
	return &Env{Vars: envVars}
}
