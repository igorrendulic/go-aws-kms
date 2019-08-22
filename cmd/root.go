package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "ckms",
	Short:            "Encrypting, decrypting configuration files with AWS KMS service",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		// help only
	},
}

var sourceFile string
var destionationFile string
var awsKmsKeyID string
var secretName string

func init() {
	rootCmd.PersistentFlags().StringVarP(&awsKmsKeyID, "key-id", "k", "", "aws kms key ID")
	rootCmd.MarkPersistentFlagRequired("kms_key_id")

	encryptCmd.Flags().StringVarP(&sourceFile, "input", "i", "", "input file")
	encryptCmd.Flags().StringVarP(&destionationFile, "output", "o", "", "output file")
	encryptCmd.MarkFlagRequired("input_file")
	encryptCmd.MarkFlagRequired("output_file")

	decryptCmd.Flags().StringVarP(&sourceFile, "input", "i", "", "input file")
	decryptCmd.Flags().StringVarP(&destionationFile, "output", "o", "", "output file")
	decryptCmd.MarkFlagRequired("input_file")
	decryptCmd.MarkFlagRequired("output_file")

	secretCmd.Flags().StringVarP(&secretName, "secret", "s", "", "name of your secret")
	secretCmd.MarkFlagRequired("secret")

	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
	rootCmd.AddCommand(secretCmd)
}

// Execute command line
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
