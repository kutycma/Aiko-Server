package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var targetVersion string

var (
	updateCommand = cobra.Command{
		Use:   "update",
		Short: "Update Aiko-Server version",
		Run: func(_ *cobra.Command, _ []string) {
			execCommandStd("bash",
				"<(curl -Ls https://raw.githubusercontents.com/Github-Aiko/Aiko-Server-script/master/install.sh)",
				targetVersion)
		},
		Args: cobra.NoArgs,
	}
	uninstallCommand = cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall Aiko-Server",
		Run:   uninstallHandle,
	}
)

func init() {
	updateCommand.PersistentFlags().StringVar(&targetVersion, "version", "", "update target version")
	command.AddCommand(&updateCommand)
	command.AddCommand(&uninstallCommand)
}

func uninstallHandle(_ *cobra.Command, _ []string) {
	var yes string
	fmt.Println(Warn("Are you sure you want to uninstall Aiko-Server? (Y/n)"))
	fmt.Scan(&yes)
	if strings.ToLower(yes) != "y" {
		fmt.Println("Uninstall canceled")
	}
	_, err := execCommand("systemctl stop Aiko-Server&&systemctl disable Aiko-Server")
	if err != nil {
		fmt.Println(Err("exec cmd error: ", err))
		fmt.Println(Err("Uninstall failed"))
		return
	}
	_ = os.RemoveAll("/etc/systemd/system/Aiko-Server.service")
	_ = os.RemoveAll("/etc/Aiko-Server/")
	_ = os.RemoveAll("/usr/local/Aiko-Server/")
	_, err = execCommand("systemctl daemon-reload&&systemctl reset-failed")
	if err != nil {
		fmt.Println(Err("exec cmd error: ", err))
		fmt.Println(Err("Uninstall failed"))
		return
	}
	fmt.Println(Ok("Uninstall successful"))
}