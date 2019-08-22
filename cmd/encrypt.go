package cmd

import (
	"fmt"
	"os"

	"github.com/igorrendulic/go-aws-kms/ckms"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Example: "ckms encrypt -i input_file.yaml -o encrypted_config.yaml -k arn:aws:kms:eu-west-2:5555555555:key/34532423-444-45aa-aaaa",
	Short:   "Encrypting configuration files with AWS KMS service",
	Run: func(cmd *cobra.Command, args []string) {
		// test
		ce := ckms.NewCKMS(awsKmsKeyID)
		fileEncrypted, err := ce.Encrypt(sourceFile)
		if err != nil {
			fmt.Printf("failed to encrypt file: %s\nError: %v\n", sourceFile, err)
			os.Exit(1)
		}
		err = ckms.SaveFile(fileEncrypted, destionationFile)
		if err != nil {
			fmt.Printf("failed to save file: %s\nError: %v\n", destionationFile, err)
			os.Exit(1)
		}
		fmt.Printf("Success! Encrypted file saved to: %s\n", destionationFile)
	},
}
