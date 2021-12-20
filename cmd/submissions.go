/*a
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/PrashantBtkl/unbundle-reddit-cli/utils"
	"github.com/spf13/cobra"
)

var subReddit, query, sortType *string

// submissionsCmd represents the submissions command
var submissionsCmd = &cobra.Command{
	Use:   "submissions",
	Short: "get reddit submissions",
	Long:  "get reddit submissions",
	Run: func(cmd *cobra.Command, args []string) {
		getSubmissions(subReddit, query, sortType)
	},
}

func init() {
	rootCmd.AddCommand(submissionsCmd)

	query = submissionsCmd.Flags().StringP("query", "q", "", "Search term. Will search ALL possible fields")
	sortType = submissionsCmd.Flags().StringP("sorttype", "", "", "Sort by a specific attribute : 'score', 'num_comments', 'created_utc'")
	subReddit = submissionsCmd.Flags().StringP("subreddit", "", "", "Restrict to a specific subreddit")
}

func getSubmissions(subReddit, query, sortType *string) {
	req, err := http.NewRequest("GET", "https://api.pushshift.io/reddit/search/submission/", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("subreddit", *subReddit)
	q.Add("query", *query)
	q.Add("sort_type", *sortType)
	req.URL.RawQuery = q.Encode()

	resp, err := utils.MakeRequest(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	fmt.Println(string(resp))

}
