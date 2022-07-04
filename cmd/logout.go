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
	wipeOk := WipeConfig()
	if wipeOk {
		color.Green("Logout successfully. All configurations are wiped out.")
	}
}

func WipeConfig() (ok bool) {
	if !IsLogin() {
		ok = false
		return
	}
	if LeetcodeClient != nil {
		LeetcodeClient = nil
	}
	// remove all config
	err := Unset("username", "password", "endpoint", "reponame", "repousername", "branch", "email", "token")
	if err != nil {
		color.Red("Failed to wipe configuration: %s", err)
		ok = false
		return
	}
	ok = true
	return
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
