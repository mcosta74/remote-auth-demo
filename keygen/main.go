package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	version *string
	purpose *string
	outDir  *string
)

func main() {

	fs := flag.NewFlagSet("key-gen", flag.ExitOnError)
	_ = parseArgs(fs)

	kg := NewKeyGen(*version)
	if kg == nil {
		fmt.Printf("no key generator for version %q\n", *version)
		os.Exit(1)
	}

	switch *purpose {
	case "public":
		generatePublic(kg, *outDir)

	case "local":
		generateLocal(kg, *outDir)

	default:
		fmt.Printf("invalid purpose %q\n", *purpose)
		os.Exit(1)
	}
	fmt.Println("Hello World")
}

func parseArgs(fs *flag.FlagSet) error {
	version = fs.String("version", "v4", "protocol version (v2, v3, v4)")
	purpose = fs.String("purpose", "", "purpose (local, public)")
	outDir = fs.String("out-dir", ".", "directory where to store generated key(s)")

	return fs.Parse(os.Args[1:])
}

func generateLocal(kg KeyGen, outDir string) {
	secret := kg.NewSymmetricKey()

	fileName := filepath.Join(outDir, fmt.Sprintf("%s-symmetric.key", kg.Version()))
	_ = os.WriteFile(fileName, []byte(secret), 0644)
}

func generatePublic(kg KeyGen, outDir string) {
	priv, pub := kg.NewAsymmetricKey()

	privFileName := filepath.Join(outDir, fmt.Sprintf("%s-asymmetric-secret.key", kg.Version()))
	pubFileName := filepath.Join(outDir, fmt.Sprintf("%s-asymmetric-public.key", kg.Version()))

	_ = os.WriteFile(privFileName, []byte(priv), 0644)
	_ = os.WriteFile(pubFileName, []byte(pub), 0644)
}
