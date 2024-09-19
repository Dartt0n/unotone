package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Addr string `koanf:"addr" validate:"required" json:"addr"`
}

func main() {
	conf := Config{Addr: "localhost:8080"}

	k := koanf.New(".")
	k.Load(env.Provider("UNOTONE_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "UNOTONE_")
		s = strings.ToLower(s)
		return s
	}), nil)
	k.Unmarshal("", &conf)

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(&conf); err != nil {
		fmt.Printf("failed ot validate config: %v\n", err)
		os.Exit(1)
	}

	resp, err := http.Get(fmt.Sprintf("http://%s/health", conf.Addr))
	if err != nil {
		fmt.Printf("failed to get healthcheck: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("health check failed with status code %d\n", resp.StatusCode)
		os.Exit(1)
	}

	os.Exit(0)
}
