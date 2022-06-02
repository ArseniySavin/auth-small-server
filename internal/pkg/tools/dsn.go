package tools

import "fmt"

func GetDsn(login, pass, host, name, sslmode string, port int) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		login, pass, host, port, name, sslmode)
}
