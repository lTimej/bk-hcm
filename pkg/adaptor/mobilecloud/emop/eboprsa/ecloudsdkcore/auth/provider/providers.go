package provider

import (
	"fmt"
	"os"
	"strings"

	"hcm/pkg/adaptor/mobilecloud/emop/eboprsa/ecloudsdkcore/auth"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/errs"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
	"gopkg.in/ini.v1"
)

type ICredentialProvider interface {
	GetCredential() (*auth.Credential, error)
}

type BasicCredentialProvider struct {
	credential *auth.Credential
}

func NewBasicCredentialProvider(credential *auth.Credential) *BasicCredentialProvider {
	if utils.IsUnSet(credential) {
		return nil
	}
	if utils.IsUnSet(credential.CredentialType) {
		credential.CredentialType = auth.CredentialTypePointer(auth.CredentialAkSk)
	}
	if utils.IsUnSet(credential.EncryptionType) {
		credential.EncryptionType = auth.EncryptionTypePointer(auth.EncrytNone)
	}
	return &BasicCredentialProvider{
		credential: credential,
	}
}

func (p *BasicCredentialProvider) GetCredential() (*auth.Credential, error) {
	return p.credential, nil
}

func CreateBasicCredentialProvider(credential *auth.Credential) *BasicCredentialProvider {
	return NewBasicCredentialProvider(credential)
}

type EnvCredentialProvider struct {
}

func NewEnvCredentialProvider() *EnvCredentialProvider {
	return &EnvCredentialProvider{}
}

func (p *EnvCredentialProvider) GetCredential() (*auth.Credential, error) {
	accessKey := os.Getenv(auth.EnvAkKey)
	secretKey := os.Getenv(auth.EnvSkKey)
	privateKey := os.Getenv(auth.EnvMopPrivateKey)
	publicKey := os.Getenv(auth.EnvMopPublicKey)
	if utils.IsSet(accessKey) && utils.IsSet(secretKey) {
		return auth.NewCredentialBuilder().
			AccessKey(accessKey).
			SecretKey(secretKey).
			PrivateKey(privateKey).
			PublicKey(publicKey).
			Build(), nil
	}
	return nil, errs.NewCredentialError("EnvCredentialProvider: accessKey or secretKey cannot be empty", nil)
}

func CreateEnvCredentialProvider() *EnvCredentialProvider {
	return NewEnvCredentialProvider()
}

type ProfileCredentialProvider struct {
	CredentialType auth.CredentialType
}

func NewProfileCredentialProvider(credentialType auth.CredentialType) *ProfileCredentialProvider {
	return &ProfileCredentialProvider{CredentialType: credentialType}
}

func getProfilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", errs.NewCredentialError("Get the user home directory error", err)
	}
	dir = strings.Replace(auth.PathCredentialFile, "~", dir, 1)
	_, err = os.Stat(dir)
	if err != nil {
		return "", nil
	}
	return dir, nil
}

func (p *ProfileCredentialProvider) GetCredential() (*auth.Credential, error) {
	path, ok := os.LookupEnv(auth.EnvCredentialFile)
	if !ok {
		defaultPath, err := getProfilePath()
		if err != nil {
			return nil, err
		}
		path = defaultPath
		if path == "" {
			return nil, nil
		}
	} else if path == "" {
		return nil, errs.NewCredentialError(auth.EnvCredentialFile+" cannot be empty", nil)
	}

	file, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	section := file.Section(string(p.CredentialType))
	if section == nil {
		return nil, errs.NewCredentialError(fmt.Sprintf("credential type '%s' does not exist in '%s'", p.CredentialType, path), nil)
	}
	if section.Key(auth.ProfileAkKey).String() == "" || section.Key(auth.ProfileSkKey).String() == "" {
		return nil, errs.NewCredentialError("accessKey or secretKey cannot be empty", nil)
	}
	return auth.NewCredentialBuilder().
		AccessKey(section.Key(auth.ProfileAkKey).String()).
		SecretKey(section.Key(auth.ProfileSkKey).String()).
		PrivateKey(section.Key(auth.ProfileMopPrivateKey).String()).
		PublicKey(section.Key(auth.ProfileMopPublicKey).String()).
		Build(), nil
}

func CreateProfileCredentialProvider(credentialType auth.CredentialType) *ProfileCredentialProvider {
	return NewProfileCredentialProvider(credentialType)
}

type CredentialProviderChain struct {
	providers []ICredentialProvider
}

func NewCredentialProviderChain(providers ...ICredentialProvider) *CredentialProviderChain {
	return &CredentialProviderChain{providers: providers}
}

func (c *CredentialProviderChain) GetCredential() (*auth.Credential, error) {
	for _, provider := range c.providers {
		credential, err := provider.GetCredential()
		if err == nil && credential != nil {
			return credential, nil
		}
	}
	return nil, errs.NewCredentialError("no valid credential found", nil)
}

func CreateCredentialProviderChain(providers ...ICredentialProvider) *CredentialProviderChain {
	return NewCredentialProviderChain(providers...)
}
