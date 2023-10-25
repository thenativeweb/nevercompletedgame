package environment

import (
	"os"
	"strconv"
)

func GetPort(defaultPort int) (int, error) {
	portAsString := os.Getenv("PORT")

	if portAsString == "" {
		return defaultPort, nil
	}

	port, err := strconv.Atoi(portAsString)
	if err != nil {
		return 0, err
	}

	return port, nil
}
