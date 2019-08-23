# Encrypting and Decrypting files with Amazon Key Service and storing secrets with Amazon Secret Manager

[![Build Status](https://travis-ci.org/igorrendulic/go-aws-kms.svg?branch=master)](https://travis-ci.org/igorrendulic/go-aws-kms)

Easy to use CLI tool for encrypting and decrypting files using AWS managed keys. It also supports Amazon Secret Manager for retrieval of secrets such as OAuth tokens, database credentials, API keys, ... 

It's intention is to enable developers to store and deploy configuration files and access credentials or other secrets stored within ASM (Amazon Secret Manager). 

[Amazon Secret Manager](https://aws.amazon.com/secrets-manager/)

[Amazon Key Service](https://aws.amazon.com/kms/)

Rotations of the keys is currently not supported.

# Prerequisite

## Create KMS Policy in AWS Console

[Go to IAM home](https://console.aws.amazon.com/iam/home)

Click Policies -> Create Policy

For service select KMS and specifiy allowed operations. Name your policy `kms-access-policy`

Create another policy for Secrets Manager (previously KMS) and specify allowed operations. For this policy read operation is enough. Name your policy `kms-secret-policy`

## Create new IAM User

Head over to [Identity and Access Management]( https://console.aws.amazon.com/iam/home). 

Click Users -> Add User. Be sure to select programatic access only.

Follow the prompts and attach existing policies directly : `kms-access-policy` and `kms-secret-policy`. 

Be sure to copy your Access Key Id and Secret access key on the final screen before proceeding with the next steps as you'll need these later.

Name your user: `kms-user`

## Create AWS KMS Key

Head over to [AWS Key Management Service](https://console.aws.amazon.com/kms). 

Click Custom managed keys -> Create Key

In the section Define key administrative permissions select user `kms-user`. 

Name the key `configkey`.

Check the created arn. You'll need it for running `ckms`. 

## Create a Secret

Head over to [Secrets Manager](https://console.aws.amazon.com/secretsmanager). 

Click Secrets -> Store new secret

Select other type of secrets and add a few secret key/values. From the dropdown select KMS key `configkey`. 

Name the secret `mysecrets`

## Use aws CLI to set credentials

```
aws configure
```
Use your users `Access Key ID` and `Secret` for credentials. 

# Installing

Using CKMS is easy. First use `go get` to install the latest version of the library. This command will install the CKMS executable along with the library and its dependencies: 

```
go get -u github.com/igorrendulic/go-aws-kms
```
Next, include CKSM in your application:
``` go
import "github.com/igorrendulic/go-aws-kms"
```

# Commands

- encrypt
- decrypt
- secret

## Command Usage

```
Encrypting, decrypting configuration files with AWS KMS service

Usage:
  ckms [flags]
  ckms [command]

Available Commands:
  decrypt     Decrypting files with AWS KMS service
  encrypt     Encrypting files with AWS KMS service
  help        Help about any command
  secret      Retrieve secret from AWS Secret Manager

Flags:
  -h, --help            help for ckms
  -k, --key-id string   aws kms key ID

Use "ckms [command] --help" for more information about a command.
```

## Flags

-k arn of your KMS key

-i input file (when encrypting location of plaintext file, when decrypting location of ciphertext file)

-o output file (when encrypting cihpertext file locaton,when decrypting location of plaintext file)

## Encrypt file

Replace placeholder ARN key with yours. 

```
ckms encrypt -i inputfile -o outputfile -k arn:aws:kms:eu-west-2:111111111:key/1111111-11111-11aa-aa11-111111
```

## Decrypt file

Replace placeholder ARN key with yours. 

```
ckms decrypt -i encrypted.yaml -o plaintest.yaml -k arn:aws:kms:eu-west-2:111111111:key/1111111-11111-11aa-aa11-111111
```

## Get secret

Replace placeholder ARN key with yours. 

```
ckms secret -s mysecrets -k arn:aws:kms:eu-west-2:111111111:key/1111111-11111-11aa-aa11-111111
```

# Production deployment

Do not store AWS API keys onto production environment. The correct way to do this is to deploy you VM instances with roles. 

Recommended resources:

[IAM Roles for Amazon EC2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html)

[Attach an AWS IAM Role to an existing EC2 Instance by using AWS CLI](https://aws.amazon.com/blogs/security/new-attach-an-aws-iam-role-to-an-existing-amazon-ec2-instance-by-using-the-aws-cli/)

[AWS IAM Integration for Kubernetes](https://kubernetes-on-aws.readthedocs.io/en/latest/user-guide/iam-roles.html)


# Getting started

Initiate CKMS library: 

``` go
kmsKey := "arn:..."
ced := NewCKMS(kmsKey)
```

Encrypt a file:
``` go
encrypted, err := ced.Ecrypt("path/to/file")
```

Decrypt a file:
```go
decryptedBytes, err := ced.Decrypt("/path/to/encrypted/file")
```

Retrieve secret:
```go
secrets, err := ced.GetSecret("mysecrets")
```

The returned result of secrets is a `map`:
```go
map[string]string
```

# Limitations

AWS KMS can encrypt only files up to 4KB (4096 bytes). These operations are designed to encrypt and decrypt data keys. Although you might use them to encrypt small amounts of data, such as a password or RSA key, they are not designed to encrypt application data.