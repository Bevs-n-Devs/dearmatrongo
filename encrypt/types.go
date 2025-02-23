package encrypt

import (
	"fmt"
	"os"

	"github.com/Bevs-n-Devs/dearmatrongo/env"
	"github.com/Bevs-n-Devs/dearmatrongo/logs"
)

var (
	masterKeyStr string // get the MASTER_KEY string from envrionment variable
	masterKey    []byte // convert the MASTER_KEY into 32 bytes for encrytion & decryption process
)

func getMasterKeyStr() error {
	if os.Getenv("MASTER_KEY") == "" {
		logs.Logs(2, "Could not get MASTER_KEY from Heroku. Loading from .env file...")
		err := env.LoadEnv("env/.env")
		if err != nil {
			logs.Logs(3, fmt.Sprintf("Unable to load environment variables: %s", err.Error()))
			return err
		}
	}

	masterKeyStr = os.Getenv("MASTER_KEY")
	if masterKeyStr == "" {
		logs.Logs(3, "MASTER_KEY is empty!")
		return fmt.Errorf("MASTER_KEY is empty")
	}

	masterKey = []byte(masterKeyStr)[:32]
	return nil
}

func InitEncryption() error {
	logs.Logs(1, "Initializing encryption functions...")
	err := getMasterKeyStr()
	if err != nil {
		logs.Logs(3, fmt.Sprintf("Error getting master key: %s", err.Error()))
		return err
	}
	logs.Logs(1, "Encryption functions successfully initialized.")
	return nil
}
