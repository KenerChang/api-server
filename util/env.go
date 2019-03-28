package util

import (
	"bufio"
	"github.com/KenerChang/api-server/util/logger"
	"os"
	"strings"
)

var (
	envs = map[string]string{}
)

func LoadConfigFromFile(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		pair := strings.Split(line, "=")

		if len(pair) == 2 {
			envs[pair[0]] = pair[1]
		}
	}

	logger.Info.Printf(nil, "envs: %+v", envs)

	err = scanner.Err()
	return
}

func LoadConfigFromEnv() (err error) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")

		if len(pair) == 2 {
			envs[pair[0]] = pair[1]
		}
	}
	return nil
}

func Get(name string) string {
	value := envs[name]
	return value
}
