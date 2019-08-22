package cmd

import (
	"fmt"
	"os"

	"github.com/igorrendulic/go-aws-kms/ckms"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Example: "ckms decrypt -i ecrypted_file.yaml -o decrypted.yaml -k arn:aws:kms:eu-west-2:5555555555:key/34532423-444-45aa-aaaa",
	Short:   "Decrypting configuration files with AWS KMS service",
	Run: func(cmd *cobra.Command, args []string) {
		// test
		ce := ckms.NewCKMS(awsKmsKeyID)
		fileDecrypted, err := ce.Decrypt(sourceFile)
		if err != nil {
			fmt.Printf("failed to decrypt file: %s\nError: %v\n", sourceFile, err)
			os.Exit(1)
		}
		err = ckms.SaveFile(fileDecrypted, destionationFile)
		if err != nil {
			fmt.Printf("failed to save file: %s\nError: %v\n", destionationFile, err)
			os.Exit(1)
		}
		fmt.Printf("Success! Decrypted file saved to: %s\n", destionationFile)
	},
}
