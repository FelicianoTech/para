package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/felicianotech/sonar/sonar/docker"
	"github.com/spf13/cobra"
)

var (
	brewFlag string
	ghFlag   string
	dhFlag   string

	// fetchCmd represents the fetch command
	fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch install data from Brew, GitHub, and/or Docker Hub",
		Run: func(cmd *cobra.Command, args []string) {

			if brewFlag != "" {
				fetchBrewData(brewFlag)
				fmt.Println("")
			}

			if ghFlag != "" {
				fetchGitHubData(ghFlag)
			}

			if dhFlag != "" {
				fetchDockerHubData(dhFlag)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.PersistentFlags().StringVar(&brewFlag, "brew", "", "Brew formula name")
	fetchCmd.PersistentFlags().StringVar(&ghFlag, "github", "", "GitHub orgname/reponame")
	fetchCmd.PersistentFlags().StringVar(&dhFlag, "dockerhub", "", "Docker image")
}

type formulaMetric struct {
	Number  int
	Formula string
	Count   string
	Percent string
}

type brewResp struct {
	Category   string          `json:"category"`
	TotalItems int             `json:"total_items"`
	StartDate  string          `json:"start_date"`
	EndDate    string          `json:"end_date"`
	TotalCount int             `json:"total_count"`
	Items      []formulaMetric `json:"items"`
}

func fetchBrewData(formula string) {

	resp, err := http.Get("https://formulae.brew.sh/api/analytics/install-on-request/30d.json")
	if err != nil {
		log.Fatal("Error: Failed to fetch Brew data.")
	}
	defer resp.Body.Close()

	jsonText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: Failed to parse Brew JSON.")
	}

	var myBrewResp brewResp

	err = json.Unmarshal(jsonText, &myBrewResp)

	fmt.Println("Brew data:")

	for _, item := range myBrewResp.Items {

		if item.Formula == formula {
			fmt.Printf("The number of installs in the last 30 days is: %s\n", item.Count)
		}
	}

}

type githubReleaseAsset struct {
	Name          string `json:"name"`
	DownloadCount int    `json:"download_count"`
}

type githubRelease struct {
	Name   string
	Assets []githubReleaseAsset
}

func fetchGitHubData(repo string) {

	resp, err := http.Get("https://api.github.com/repos/" + repo + "/releases")
	if err != nil {
		log.Fatal("Error: Failed to fetch GitHub data.")
	}
	defer resp.Body.Close()

	jsonText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: Failed to parse GitHub JSON.")
	}

	var ghReleases []githubRelease

	err = json.Unmarshal(jsonText, &ghReleases)

	fmt.Println("GitHub data:")

	for _, release := range ghReleases {

		fmt.Printf("%s:\n", release.Name)

		for _, asset := range release.Assets {

			fmt.Printf("%s: %d\n", asset.Name, asset.DownloadCount)
		}
	}

}

func fetchDockerHubData(image string) {

	pulls, err := docker.ImagePulls(image)
	if err != nil {
		log.Fatal("Error: Failed to fetch Docker Hub pulls.")
	}

	stars, err := docker.ImageStars(image)
	if err != nil {
		log.Fatal("Error: Failed to fetch Docker Hub stars.")
	}

	fmt.Println("Docker Hub data:")
	fmt.Printf("The number of pulls for %s is: %d\n", image, pulls)
	fmt.Printf("The number of stars for %s is: %d\n", image, stars)
}
