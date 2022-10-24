package defenderclienthttp

import (
	kilncmdutils "github.com/kilnfi/go-utils/cmd/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Flags register viper compatible pflags for auth
func Flags(v *viper.Viper, f *pflag.FlagSet) {
	UserPoolIDFlag(v, f)
	ClientPoolIDFlag(v, f)
	APIKeyFlag(v, f)
	APISecretFlag(v, f)
	APIURLFlag(v, f)
}

const (
	userPoolIDFlag     = "defender-user-pool-id"
	userPoolIDViperKey = "defender.userPoolID"
	UserPoolIDEnv      = "DEFENDER_USER_POOL_ID"
	UserPoolIDDefault  = "us-west-2_94f3puJWv"
)

// UserPoolIDFlag register flag for Authentication userPoolID
func UserPoolIDFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Defender cognito user pool ID",
		UserPoolIDEnv,
	)

	f.String(userPoolIDFlag, "", desc)
	_ = v.BindPFlag(userPoolIDViperKey, f.Lookup(userPoolIDFlag))
	_ = v.BindEnv(userPoolIDViperKey, UserPoolIDEnv)
}

func GetUserPoolID(v *viper.Viper) string {
	return v.GetString(userPoolIDViperKey)
}

const (
	clientPoolIDFlag     = "defender-client-pool-id"
	clientPoolIDViperKey = "defender.clientPoolID"
	ClientPoolIDEnv      = "DEFENDER_CLIENT_POOL_ID"
	ClientPoolIDDefault  = "40e58hbc7pktmnp9i26hh5nsav"
)

// ClientPoolIDFlag register flag for Authentication clientPoolID
func ClientPoolIDFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Defender cognito client pool ID",
		ClientPoolIDEnv,
	)

	f.String(clientPoolIDFlag, "", desc)
	_ = v.BindPFlag(clientPoolIDViperKey, f.Lookup(clientPoolIDFlag))
	_ = v.BindEnv(clientPoolIDViperKey, ClientPoolIDEnv)
}

func GetClientPoolID(v *viper.Viper) string {
	return v.GetString(clientPoolIDViperKey)
}

const (
	apiKeyFlag     = "defender-api-key"
	apiKeyViperKey = "defender.apiKey"
	APIKeyEnv      = "DEFENDER_API_KEY"
)

// APIKeyFlag register flag for Authentication apiKey
func APIKeyFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Defender API Key",
		APIKeyEnv,
	)

	f.String(apiKeyFlag, "", desc)
	_ = v.BindPFlag(apiKeyViperKey, f.Lookup(apiKeyFlag))
	_ = v.BindEnv(apiKeyViperKey, APIKeyEnv)
}

func GetAPIKey(v *viper.Viper) string {
	return v.GetString(apiKeyViperKey)
}

const (
	apiSecretFlag     = "defender-api-secret"
	apiSecretViperKey = "defender.apiSecret"
	APISecretEnv      = "DEFENDER_API_SECRET"
)

// APISecretFlag register flag for Authentication apiSecret
func APISecretFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Defender API Secret",
		APISecretEnv,
	)

	f.String(apiSecretFlag, "", desc)
	_ = v.BindPFlag(apiSecretViperKey, f.Lookup(apiSecretFlag))
	_ = v.BindEnv(apiSecretViperKey, APISecretEnv)
}

func GetAPISecret(v *viper.Viper) string {
	return v.GetString(apiSecretViperKey)
}

const (
	apiURLFlag     = "defender-api-url"
	apiURLViperKey = "defender.apiURL"
	APIURLEnv      = "DEFENDER_API_URL"
	APIURLDefault  = "https://defender-api.openzeppelin.com/admin"
)

// APIURLFlag register flag for Authentication apiURL
func APIURLFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Defender API URL",
		APIURLEnv,
	)

	f.String(apiURLFlag, "", desc)
	_ = v.BindPFlag(apiURLViperKey, f.Lookup(apiURLFlag))
	_ = v.BindEnv(apiURLViperKey, APIURLEnv)
}

func GetAPIURL(v *viper.Viper) string {
	return v.GetString(apiURLViperKey)
}

func NewConfigFromViper(v *viper.Viper) *Config {
	return &Config{
		APIKey:       GetAPIKey(v),
		APISecret:    GetAPISecret(v),
		APIURL:       GetAPIURL(v),
		UserPoolID:   GetUserPoolID(v),
		ClientPoolID: GetClientPoolID(v),
	}
}
