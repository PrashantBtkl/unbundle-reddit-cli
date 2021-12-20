/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var subReddit, query, sortType *string

// submissionsCmd represents the submissions command
var submissionsCmd = &cobra.Command{
	Use:   "submissions",
	Short: "get reddit submissions",
	Long:  "get reddit submissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("submissions called", *subReddit, *query, *sortType)
		PushShiftAPI(subReddit, query, sortType)
	},
}

func init() {
	rootCmd.AddCommand(submissionsCmd)

	query = submissionsCmd.Flags().StringP("query", "q", "", "Search term. Will search ALL possible fields")
	sortType = submissionsCmd.Flags().StringP("sorttype", "", "", "Sort by a specific attribute : 'score', 'num_comments', 'created_utc'")
	subReddit = submissionsCmd.Flags().StringP("subreddit", "", "", "Restrict to a specific subreddit")
}

func PushShiftAPI(subReddit, query, sortType *string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.pushshift.io/reddit/search/", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("subreddit", *subReddit)
	q.Add("query", *query)
	q.Add("sorttype", *sortType)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
