package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd represents the verify command
var runCmd = &cobra.Command{
	Use:   "check <name>",
	Short: "Check if a name is available on Snap, Brew, etc",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("General")
		fmt.Println("==========")
		resp, err := http.Get("https://formulae.brew.sh/api/formula/" + args[0] + ".json")
		if err != nil {
			log.Error(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			fmt.Println("Brew: available")
		} else if resp.StatusCode == 200 {
			fmt.Println("Brew: unavailable")
		} else {
			fmt.Println("Brew: not sure")
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.snapcraft.io/v2/snaps/info/"+args[0], nil)
		req.Header.Add("Snap-Device-Series", "16")
		resp, err = client.Do(req)
		if err != nil {
			log.Error(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			fmt.Println("Snap: available")
		} else if resp.StatusCode == 200 {
			fmt.Println("Snap: unavailable")
		} else {
			fmt.Println("Snap: not sure")
		}

		resp, err = http.Get("https://chocolatey.org/packages/" + args[0])
		if err != nil {
			log.Error(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			fmt.Println("Chocolatey: available")
		} else if resp.StatusCode == 200 {
			fmt.Println("Chocolatey: unavailable")
		} else {
			fmt.Println("Chocolatey: not sure")
		}

		fmt.Println("\nLanguage Specific")
		fmt.Println("==========")

		resp, err = http.Get("https://www.npmjs.com/package/" + args[0])
		if err != nil {
			log.Error(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			fmt.Println("NPM: available")
		} else if resp.StatusCode == 200 {
			fmt.Println("NPM: unavailable")
		} else {
			fmt.Println("NPM: not sure")
		}
	},
}

func init() {

	rootCmd.AddCommand(runCmd)
}
