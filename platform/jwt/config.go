package jwt

type Config struct {
	SigningKey string `yaml:"signing_key,omitempty"`
	SaltHash   int    `yaml:"salt_hash,omitempty"`
	Expiry     int    `yaml:"expiry,omitempty"`
}
