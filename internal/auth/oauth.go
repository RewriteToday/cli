package auth

import (
	"fmt"
	"time"
)

func RunOAuthFlow() (string, error) {
	fmt.Println("Opening browser for authentication...")
	fmt.Println("(Mock: skipping browser, generating demo key)")
	time.Sleep(time.Second)
	return fmt.Sprintf("rw_live_mock_%d", time.Now().UnixNano()%100000), nil
}
