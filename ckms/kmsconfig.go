package ckms

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// CKMS reads the encrypted file, decrypts it and returns
type CKMS struct {
	svc   *kms.KMS
	sm    *secretsmanager.SecretsManager
	keyID string
}

// NewCKMS - config init with aws session from the shared credentials file ~/.aws/credentials and configuration from the shared configuration file ~/.aws/config.
func NewCKMS(keyID string) *CKMS {
	if keyID == "" {
		log.Fatal("missing keyID")
	}
	regionTmp := strings.Split(keyID, "arn:aws:kms:")
	region := "us-west-1"
	if len(regionTmp) > 1 {
		index := strings.Index(regionTmp[1], ":")
		if index >= 0 {
			region = regionTmp[1][:index]
		}
	}
	// Initialize a session that the SDK uses to load
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	if err != nil {
		log.Fatalf("failed to initialize session:%v\n", err)
	}
	// Create KMS service client
	svc := kms.New(sess)
	sm := secretsmanager.New(sess)

	return &CKMS{
		svc:   svc,
		sm:    sm,
		keyID: keyID,
	}
}

// EncryptRaw converts raw bytes to base64 format and encrypts the base64 string. result is encrypted byte[]
func (ckms *CKMS) EncryptRaw(raw []byte) ([]byte, error) {
	// encodedContent := base64.StdEncoding.EncodeToString(raw)
	// Encrypt the data
	result, err := ckms.svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(ckms.keyID),
		Plaintext: raw,
	})
	if err != nil {
		return nil, err
	}
	return result.CiphertextBlob, nil
}

// Encrypt encrypts the file contents of the given input file and returns encrypted file contents in base64 format
func (ckms *CKMS) Encrypt(path string) ([]byte, error) {
	content, err := loadFile(path)

	if err != nil {
		return nil, err
	}
	result, err := ckms.EncryptRaw(content)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptRaw decrypts raw bytes (base64 formatted in encryption process)
func (ckms *CKMS) DecryptRaw(cipherText []byte) ([]byte, error) {

	result, err := ckms.svc.Decrypt(&kms.DecryptInput{CiphertextBlob: cipherText})
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}

// Decrypt the base64 config file and return plaintext contents of a file
func (ckms *CKMS) Decrypt(path string) ([]byte, error) {
	content, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	result, err := ckms.DecryptRaw(content)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSecret returns a JSON file with stored secrets in AWS Secret Manager
func (ckms *CKMS) GetSecret(secretName string) (map[string]string, error) {
	output, err := ckms.sm.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
	if err != nil {
		panic(err.Error())
	}
	var secretObj map[string]string
	err = json.Unmarshal([]byte(*output.SecretString), &secretObj)
	if err != nil {
		return nil, err
	}

	return secretObj, nil
}
