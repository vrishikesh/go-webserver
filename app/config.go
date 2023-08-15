package app

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var Config *config

type config struct {
	AppEnv          string `env:"APP_ENV" defaultEnv:"development"`
	AppHost         string `env:"APP_HOST"`
	SomethingRandom string `env:"SOMETHING_RANDOM" defaultEnv:"random"`
	DbHost          string `env:"DB_HOST"`
	DbPort          int    `env:"DB_PORT"`
	DbUser          string `env:"DB_USER"`
	DbPass          string `env:"DB_PASS"`
}

func NewConfig(embedFS embed.FS, logger *log.Logger) (*config, error) {
	return ParseConfig(embedFS, logger)
}

func ParseConfig(embedFS embed.FS, logger *log.Logger) (*config, error) {
	var config config
	var m = make(map[string]string)

	file, err := embedFS.Open("configs/development.env")
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		ss := strings.Split(text, "=")
		m[ss[0]] = ss[1]
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not scan env: %s", err)
	}

	st := reflect.TypeOf(config)
	stv := reflect.ValueOf(&config).Elem()

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		f := stv.Field(i)

		key := field.Tag.Get("env")
		defaultEnv := field.Tag.Get("defaultEnv")

		value := m[key]
		if value == "" {
			value = defaultEnv
		}
		if value == "" {
			return nil, fmt.Errorf("could not read config key %s", key)
		}

		switch {
		case f.Kind() == reflect.String && f.CanSet():
			f.SetString(value)
		case f.Kind() == reflect.Int && f.CanSet():
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("could not covert string to int: %s", value)
			}
			f.SetInt(i)
		default:
			logger.Fatalf("could not set key for %s", key)
		}
	}

	return &config, nil
}
