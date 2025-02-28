package main

import (
	"fmt"

	clientServer "github.com/aczietlow/auth-sandbox/pkg/ClientServer"
	openidconnect "github.com/aczietlow/auth-sandbox/pkg/OpenIDConnect"
)

func main() {
	// Get fancy client
	// TODO: move this to the server pkg?
	oidc, err := openidconnect.New("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("New OIDC Client created: %s", oidc.ClientID)

	clientServer.Start(oidc)
}
