package main

import (
	"fmt"

	"github.com/openshift/library-go/pkg/crypto"
	"gomodules.xyz/jsonpatch/v2"
)

func main() {
	names := []string{"VersionTLS13"}
	for _, cipherName := range crypto.OpenSSLToIANACipherSuites(names) {
		fmt.Println(cipherName)
	}
	fmt.Println(jsonpatch.NewOperation("add", "/spec/template/spec/volumes", names))

}
