package cmd

import (
	"log"
  "os"
  "os/exec"
  "github.com/spf13/cobra"
)

// brewCmd represents the brew command
var brewCmd = &cobra.Command{
	Use:   "brew",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("brew called")
    if brewInstalled() {
      brewVersion()
    } else {
      installer := downloadInstaller()
      install(installer)
    }
	},
}

func brewInstalled() bool {
  log.Println("Checking if Homebrew is installed")
  err := exec.Command("which", "brew").Run()
  return err != nil // TODO: reversed for debugging
}

func brewVersion() {
  if version, err := exec.Command("brew", "--version").Output(); err != nil {
    log.Fatal("Something went wrong checking homebrew version")
  } else {
    log.Println(string(version))
  }
}

func downloadInstaller() string {
  return "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install < /dev/null)\""

  // attempt at being smarter about downloading install script
  //   log.Println("Downloading homebrew installer.")
  //   resp, err := http.Get("https://raw.githubusercontent.com/Homebrew/install/master/install")
  //   if err != nil {
  //     log.Fatal("Could not download Homebrew.", err)
  //   }
  //   defer resp.Body.Close()
  //   body, err := ioutil.ReadAll(resp.Body)
  //   return string(body)
}

func install(installer string) {
  ruby := exec.Command("/usr/bin/ruby", "-e", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install < /dev/null)\"")
  ruby.Stdout = os.Stdout
  ruby.Stdin = os.Stdin
  ruby.Stderr = os.Stderr
  err := ruby.Run()
  if err != nil {
		log.Fatal("Something went wrong while installing homebrew", err)
	} else {
    log.Println("Homebrew is now installed")
  }
}

func init() {
	rootCmd.AddCommand(brewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// brewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// brewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
