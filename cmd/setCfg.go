package cmd

import (
	"github.com/CallanBi/leetecho-cli/leetcode_client/helper"
	color "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCfgCmd = &cobra.Command{
	Use:   "setcfg [-r/--reponame]=[value] [-n/--repousername]=[value] [-b/--branch]=[value] [-e/--email]=[value] [-t/--token]=[value]",
	Short: "Set current user's and its configuration.",
	Long: `Set current user and its configuration, including reponame, repousername, branch, email and token.
repousername is optional, if not set, it will be set to username.
branch is optional, if not set, it will be set to "main".
The other fields are required.`,
	Run:     setCfgMain,
	Example: `leetecho setcfg -r=My-LeetCode-Notes -b=master --email=yourmail@example.com -t=tokenExample`,
}

var reponame, repousername, branch, email, token *string = new(string), new(string), new(string), new(string), new(string)

type CfgParams map[string]string

var params = &CfgParams{
	"reponame":     *reponame,
	"repousername": *repousername,
	"branch":       *branch,
	"email":        *email,
	"token":        *token,
}

func init() {
	RootCmd.AddCommand(setCfgCmd)
	setCfgCmd.Flags().StringVarP(reponame, "reponame", "r", "", "repository name")
	setCfgCmd.Flags().StringVarP(repousername, "repousername", "n", "", "repository owner")
	setCfgCmd.Flags().StringVarP(branch, "branch", "b", "main", "branch name")
	setCfgCmd.Flags().StringVarP(email, "email", "e", "", "email")
	setCfgCmd.Flags().StringVarP(token, "token", "t", "", "token")
}

func setCfgMain(cmd *cobra.Command, args []string) {
	if !IsLogin() {
		return
	}

	writeArgsIntoParams := func(params *CfgParams) error {
		for k := range *params {
			value, err := cmd.Flags().GetString(k)
			if err != nil {
				return err
			}
			(*params)[k] = value
		}
		return nil
	}

	writeArgsIntoParams(params)

	if err := validators.ValidateAll(params); err != nil {
		color.Red("Invalid configuration. Error information is as follows:")
		color.Red(err.Error())
		return
	}
	if *repousername == "" {
		*repousername = viper.GetString("username")
	}
	if *branch == "" {
		*branch = "main"
	}
	viper.Set("reponame", reponame)
	viper.Set("repousername", repousername)
	viper.Set("branch", branch)
	viper.Set("email", email)
	viper.Set("token", token)
	if writeCfgErr := viper.WriteConfig(); writeCfgErr != nil {
		color.Red("An error occur when writing configuration file. Error Message: %s", writeCfgErr)
		return
	}
	color.Green("Configuration set successfully.")
}

type ValidatorMap map[string]func(string) error

func (v ValidatorMap) Validate(key string, value string) error {
	if validator, ok := v[key]; ok {
		return validator(value)
	}
	return nil
}

func (v ValidatorMap) ValidateAll(m *CfgParams) error {
	for key, value := range *m {
		if err := v.Validate(key, value); err != nil {
			return err
		}
	}
	return nil
}

// func (v ValidatorMap) ValidateAllWithPrefix(prefix string, m map[string]string) error {
// 	for key, value := range m {
// 		if err := v.Validate(prefix+key, value); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (v ValidatorMap) ValidateAllWithPrefixes(prefixes []string, m map[string]string) error {
// 	for _, prefix := range prefixes {
// 		if err := v.ValidateAllWithPrefix(prefix, m); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (v ValidatorMap) ValidateAllWithPrefixesAndSeparator(prefixes []string, separator string, m map[string]string) error {
// 	for _, prefix := range prefixes {
// 		if err := v.ValidateAllWithPrefix(prefix+separator, m); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (v ValidatorMap) ValidateAllWithPrefixesAndSeparators(prefixes []string, separators []string, m map[string]string) error {
// 	for _, prefix := range prefixes {
// 		for _, separator := range separators {
// 			if err := v.ValidateAllWithPrefix(prefix+separator, m); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

var validators = &ValidatorMap{
	"reponame": func(value string) error {
		if value == "" {
			return &helper.ErrorResp{
				Code:    helper.INVALID_CONFIG_CODE,
				Message: "Repository name cannot be empty.",
				Status:  helper.INVALID_CONFIG_STATUS,
			}
		}
		return nil
	},
	"repousername": func(value string) error {
		if value == "" {
			return nil
		}
		return nil
	},
	"branch": func(value string) error {
		if value == "" {
			return nil
		}
		return nil
	},
	"email": func(value string) error {
		if value == "" {
			return &helper.ErrorResp{
				Code:    helper.INVALID_CONFIG_CODE,
				Message: "Email cannot be empty.",
				Status:  helper.INVALID_CONFIG_STATUS,
			}
		}
		return nil
	},
	"token": func(value string) error {
		if value == "" {
			return &helper.ErrorResp{
				Code:    helper.INVALID_CONFIG_CODE,
				Message: "Token cannot be empty.",
				Status:  helper.INVALID_CONFIG_STATUS,
			}
		}
		return nil
	},
}
