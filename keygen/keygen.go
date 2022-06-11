package main

import "aidanwoods.dev/go-paseto"

type KeyGen interface {
	NewSymmetricKey() string
	NewAsymmetricKey() (private, public string)
	Version() string
}

func NewKeyGen(version string) KeyGen {
	switch version {
	case "v2":
		return &v2KeyGen{}

	case "v3":
		return &v3KeyGen{}

	case "v4":
		return &v4KeyGen{}

	default:
		return nil
	}
}

type v2KeyGen struct{}

func (kg *v2KeyGen) NewSymmetricKey() string {
	key := paseto.NewV2SymmetricKey()
	return key.ExportHex()
}

func (kg *v2KeyGen) NewAsymmetricKey() (private, public string) {
	privateKey := paseto.NewV2AsymmetricSecretKey()
	return privateKey.ExportHex(), privateKey.Public().ExportHex()
}

func (kg *v2KeyGen) Version() string {
	return "v2"
}

type v3KeyGen struct{}

func (kg *v3KeyGen) NewSymmetricKey() string {
	key := paseto.NewV3SymmetricKey()
	return key.ExportHex()
}

func (kg *v3KeyGen) NewAsymmetricKey() (private string, public string) {
	privateKey := paseto.NewV3AsymmetricSecretKey()
	return privateKey.ExportHex(), privateKey.Public().ExportHex()
}

func (kg *v3KeyGen) Version() string {
	return "v3"
}

type v4KeyGen struct{}

func (kg *v4KeyGen) NewSymmetricKey() string {
	key := paseto.NewV4SymmetricKey()
	return key.ExportHex()
}

func (kg *v4KeyGen) NewAsymmetricKey() (private string, public string) {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	return privateKey.ExportHex(), privateKey.Public().ExportHex()
}

func (kg *v4KeyGen) Version() string {
	return "v4"
}
