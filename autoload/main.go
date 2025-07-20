package autoload

import godotenv "github.com/errybase/go-dotenv"

func init() {
	godotenv.Load(".env")
}
