package cmd

import (
  "log"
  "os"
  "os/exec"
	"github.com/spf13/cobra"
)

// xcodeCmd represents the xcode command
var xcodeCmd = &cobra.Command{
	Use:   "xcode",
	Short: "Install the latest version of xcode command-line tools",
	Long: `Check if xcode command-line tools are installed, and if not, download and install the latest.
Unlike traditional installation methods, does not require manual confirmation from the user.`,
	Run: func(cmd *cobra.Command, args []string) {
    if xcodeInstalled() {
      xcodeVersion()
    } else {
      createFakeUpdate()
      installUpdates()
    }
	},
}

func xcodeInstalled() bool {
  log.Println("Checking if Xcode is installed")
  err := exec.Command("xcode-select", "--print-path").Run()
  if err != nil {
    return true
  } else {
    return false
  }
}

func xcodeVersion() {
  version, err := exec.Command("xcode-select", "--version").CombinedOutput()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Currently installed Xcode:")
  log.Println(version)
}

func createFakeUpdate() {
  fakeUpdate, err := os.Create("/tmp/.com.apple.dt.CommandLineTools.installondemand.in-progress")
  if err != nil {
    log.Fatal(err)
  }
  fakeUpdate.Close()
}

func installUpdates() {
  log.Println("Installing Xcode")
  if exec.Command("softwareupdate", "-i", "-a").Run() != nil {
    log.Fatal("Software updates could not be installed.")
  }
}

func init() {
	rootCmd.AddCommand(xcodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xcodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xcodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
