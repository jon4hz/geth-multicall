package multicall

import "fmt"

type Option func(*Config)

type Config struct {
	MulticallAddress string
	Gas              string
}

const (
	MainnetAddress    = "0x5eb3fa2dfecdde21c950813c665e9364fa609bd2"
	RopstenAddress    = "0xf3ad7e31b052ff96566eedd218a823430e74b406"
	BSCAddress        = "0x6Cf63cC81660Dd174A49e0C61A1f916456Ee1471"
	BSCTestnetAddress = "0xD3c6D8dAa57dfD38609047447cccDEF7Db6631b5"
	PolygonAddress    = "0x8a233a018a2e123c0D96435CF99c8e65648b429F"
	FantomAddress     = "0x08AB4aa09F43cF2D45046870170dd75AE6FBa306"
	CronosAddress     = "0x845C4753954c347175B4179B2D5B18DE1629f94F"
)

func WithContractAddress(address string) Option {
	return func(c *Config) {
		c.MulticallAddress = address
	}
}

func WithGas(gas uint64) Option {
	return func(c *Config) {
		c.Gas = fmt.Sprintf("0x%x", gas)
	}
}

func WithGasHex(gas string) Option {
	return func(c *Config) {
		c.Gas = gas
	}
}
