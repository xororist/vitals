/*
Copyright © 2022 Hugo Cachon <hugo.cachon@hetic.net>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	LighthouseResult LightHouseResult `json:"lighthouseResult"`
}

type LightHouseResult struct {
	LightHouseVersion string `json:"lightHouseVersion"`
	RequestedUrl      string `json:"requestedUrl"`
	FetchTime         string `json:"fetchTime"`
	Categories        struct {
		Seo struct {
			ID    string  `json:"id"`
			Title string  `json:"title"`
			Score float64 `json:"score"`
		} `json:"seo"`
		BestPractices struct {
			ID    string  `json:"id"`
			Title string  `json:"title"`
			Score float64 `json:"score"`
		} `json:"best-practices"`
		Accessibility struct {
			ID    string  `json:"id"`
			Title string  `json:"title"`
			Score float64 `json:"score"`
		} `json:"accessibility"`
		Performance struct {
			ID    string  `json:"id"`
			Title string  `json:"title"`
			Score float64 `json:"score"`
		} `json:"performance"`
	} `json:"categories"`
}

var (
	urlPath string
	client  = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
)

func Request(Url string) (*http.Response, error) {
	str := "https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url="

	u, _ := url.Parse(str)
	//fmt.Println("original:", u)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("strategy", "MOBILE")
	q.Add("category", "PERFORMANCE")

	u.RawQuery = q.Encode()

	//fmt.Println("modified:", u)

	ApiKey := "AIzaSyB5HqcLm-6widAtJAzc2S-Q09DyUMvXy-g"
	req, err := http.NewRequest("GET", str+Url+"&key="+ApiKey+"&strategy=mobile&category=seo&category=best-practices&category=accessibility&category=performance", nil)
	response, err := client.Do(req)

	return response, err
}

func DataToJson(data []byte) Response {
	var responseObject Response
	err := json.Unmarshal(data, &responseObject)
	if err != nil {
		fmt.Print(err.Error())
	}
	return responseObject
}
func ping(domain string) (int, error) {
	req, err := http.NewRequest("HEAD", domain, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	err = resp.Body.Close()
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func check(domain string) {
	checkUrl := "https://" + domain
	fmt.Println(ping(checkUrl))

	_, err := ping(checkUrl)
	if err != nil {
		return
	}
	response, err := Request(checkUrl)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseObject := DataToJson(responseData)

	fmt.Printf("Lighthouse Version: %s\n", responseObject.LighthouseResult.LightHouseVersion)
	fmt.Printf("Lighthouse Fetch Time: %s\n", responseObject.LighthouseResult.FetchTime)
	fmt.Printf("Lighthouse Seo Score: %v\n", responseObject.LighthouseResult.Categories.Seo.Score)
	fmt.Printf("Lighthouse Best Practices Score: %v\n", responseObject.LighthouseResult.Categories.BestPractices.Score)
	fmt.Printf("Lighthouse Performance Score: %v\n", responseObject.LighthouseResult.Categories.Performance.Score)
	fmt.Printf("Lighthouse Accessibility Score: %v\n", responseObject.LighthouseResult.Categories.Accessibility.Score)
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run a check on a specified domain name using the -u flag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		check(urlPath)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.
	checkCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to check")
	if err := checkCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
