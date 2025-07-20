package godotenv_test

import (
	"os"
	"testing"

	godotenv "github.com/errybase/go-dotenv"
)

func TestLoad(t *testing.T) {
	if err := godotenv.Load(".env_test"); err != nil {
		t.Error(err)
	}

	if want, got := "foo", os.Getenv("FOO"); want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
	if want, got := "bar", os.Getenv("BAR"); want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
