package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	color "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear user configuration and logout",
	Long:  `Clear user configuration and logout.`,
	Run:   logoutMain,
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}

func logoutMain(cmd *cobra.Command, args []string) {
	WipeConfig()
	color.Green("Logout successfully. All configurations are wiped out.")
}

func WipeConfig() {
	if LeetcodeClient != nil {
		LeetcodeClient = nil
	}
	// remove all config
	Unset("username", "password", "endpoint", "reponame", "repousername", "branch", "email", "token")
}

func Unset(vars ...string) error {
	cfg := viper.AllSettings()
	vals := cfg

	for _, v := range vars {
		parts := strings.Split(v, ".")
		for i, k := range parts {
			v, ok := vals[k]
			if !ok {
				// Doesn't exist no action needed
				break
			}

			switch len(parts) {
			case i + 1:
				// Last part so delete.
				delete(vals, k)
			default:
				m, ok := v.(map[string]interface{})
				if !ok {
					return fmt.Errorf("unsupported type: %T for %q", v, strings.Join(parts[0:i], "."))
				}
				vals = m
			}
		}
	}

	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err = viper.ReadConfig(bytes.NewReader(b)); err != nil {
		return err
	}

	return viper.WriteConfig()
}
