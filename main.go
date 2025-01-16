package main

import (
	"fmt"

	"github.com/openshift/library-go/pkg/crypto"
)

func main() {
	names := []string{"VersionTLS13"}
	for _, cipherName := range crypto.OpenSSLToIANACipherSuites(names) {
		fmt.Println(cipherName)
	}

}
