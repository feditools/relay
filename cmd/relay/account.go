package main

import (
	"github.com/feditools/relay/cmd/relay/action/account"
	"github.com/feditools/relay/cmd/relay/flag"
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/cobra"
)

// accountCommands returns the 'account' subcommand.
func accountCommands() *cobra.Command {
	accountCmd := &cobra.Command{
		Use:   "account",
		Short: "manage accounts",
	}
	flag.Account(accountCmd, config.Defaults)

	accountModifyCmd := &cobra.Command{
		Use:   "modify",
		Short: "modify an account",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), account.Modify)
		},
	}
	flag.AccountModify(accountModifyCmd, config.Defaults)
	accountCmd.AddCommand(accountModifyCmd)

	return accountCmd
}
