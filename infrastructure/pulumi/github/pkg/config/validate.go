package config

import (
	"fmt"

	"github.com/gookit/validate"
)

func ValidateConfiguration(cfg *Config) error {
	v := validate.Struct(cfg)
	v.StopOnError = false
	if !v.Validate() {
		err := "There are issues with configuration"
		fmt.Printf(err + ":\n")
		for _, problem := range v.Errors {
			fmt.Printf("  - " + problem.One() + "\n")
		}
		return fmt.Errorf(err)
	}
	return nil
}

func (cfg Config) IsVisibility(s string) bool {
	if s != "public" && s != "private" {
		return false
	}
	return true
}

func (cfg Config) IsColorHex(s string) bool {
	if len(s) != 6 {
		return false
	}

	for _, c := range s {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F') {
			return false
		}
	}
	return true
}

func (cfg Config) Messages() map[string]string {
	return validate.MS{
		"required":     "{field} is required",
		"isVisibility": "{field} is either 'public' or 'private'",
		"isColorHex":   "{field} must be 6 character color hex code without the leading '#'",
	}
}
