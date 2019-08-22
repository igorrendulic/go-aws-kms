package ckms

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Secret1 string `yaml:"my_secret_key1"`
	Secret2 string `yaml:"my_secret_key2"`
}

// TestEncryptDecrypt requires environment variables to be set
func TestEncryptDecrypt(t *testing.T) {
	kmsKey := os.Getenv("AWS_KMS_KEY")

	ced := NewCKMS(kmsKey)

	encrypted, err := ced.Encrypt("../testresources/testconfig.yaml")
	if err != nil {
		t.Fatal(err)
	}
	rawContent, err := ced.DecryptRaw(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	var config Config
	err = yaml.Unmarshal(rawContent, &config)
	if err != nil {
		t.Fatal(err)
	}
	if config.Secret1 != "hello" {
		t.Fatal("unexpected value in testconfig.yaml")
	}
	if config.Secret2 != "goodbye" {
		t.Fatal("unexpected value in testconfig.yaml")
	}
}

func TestGetSecret(t *testing.T) {
	kmsKey := os.Getenv("AWS_KMS_KEY")

	ced := NewCKMS(kmsKey)
	secretObj, err := ced.GetSecret("unittestsecret")
	if err != nil {
		t.Fatal(err)
	}
	if "testsecretapikey1" != secretObj["apikey1"] {
		t.Fatal("apikey1 didn't match a secret")
	}
}
