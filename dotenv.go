package godotenv

import (
	"os"

	"github.com/errybase/go-dotenv/parser"
)

func Load(name string) error {
	b, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	for k, v := range parser.Parse(b) {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}

	return nil
}
