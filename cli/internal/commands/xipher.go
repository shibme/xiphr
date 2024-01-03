package commands

import "github.com/spf13/cobra"

func XipherCommand() *cobra.Command {
	if xipherCmd != nil {
		return xipherCmd
	}
	xipherCmd = &cobra.Command{
		Use:   "xipher",
		Short: "Xipher is a curated collection of cryptographic primitives written in Go to encrypt and decrypt data with optional compression.",
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetBool(versionFlag.name)
			if version {
				showVersionInfo()
			} else {
				cmd.Help()
			}
		},
	}
	xipherCmd.Flags().BoolP(versionFlag.name, versionFlag.shorthand, false, versionFlag.usage)
	xipherCmd.AddCommand(versionCommand())
	return xipherCmd
}
