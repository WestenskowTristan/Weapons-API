package env

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(BasePath(".env")) // Can ignore an error here.
}

type ConfigurationMissingError struct {
	Key string
}

func (err *ConfigurationMissingError) Error() string {
	return fmt.Sprintf("configuration-missing:%s", err.Key)
}

func BasePath(relativePaths ...string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	srcIndex := strings.LastIndex(basepath, "/src")

	paths := []string{basepath[0:srcIndex]}
	newPath := path.Join(append(paths, relativePaths...)...)
	return newPath
}

func TryGet(key string) (string, bool) {
	return os.LookupEnv(key)
}

func Get(key string) string {
	value, ok := TryGet(key)
	if !ok || value == "" {
		panic(&ConfigurationMissingError{
			Key: key,
		})
	}

	return value
}

func GetInt(key string, defaultValue *int) int {
	value, ok := TryGet(key)
	if !ok {
		if defaultValue != nil {
			return *defaultValue
		}

		panic(&ConfigurationMissingError{
			Key: key,
		})
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return i
}