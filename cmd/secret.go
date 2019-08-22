package cmd

import (
	"fmt"

	"github.com/igorrendulic/go-aws-kms/ckms"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:     "secret",
	Example: "ckms secret -s mysecretsname -k arn:aws:kms:eu-west-2:5555555555:key/34532423-444-45aa-aaaa",
	Short:   "retrieve secret from AWS Secret Manager",
	Run: func(cmd *cobra.Command, args []string) {
		// test
		ce := ckms.NewCKMS(awsKmsKeyID)
		secret, err := ce.GetSecret(secretName)
		if err != nil {
			fmt.Printf("failed to retrieve secret: %v\n", err)
			panic(err)
		}

		fmt.Printf("Your secret (key:value): \n")
		if secret == nil {
			fmt.Printf("secret empty")
		} else {
			for k, v := range secret {
				fmt.Printf("%s : %s\n", k, v)
			}
		}
	},
}
