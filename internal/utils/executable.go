package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetCwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(filepath.Dir(ex))
}
