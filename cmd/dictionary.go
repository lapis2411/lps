package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/lapis2411/lps/config"
)

const DICTIONARY_URL = "https://api.api-ninjas.com/v1/dictionary?word="

func GetDictionaryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dictionary",
		Short: "get synonym of English word",
		Long:  "get synonym of English word",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			word := args[0]
			searchDictionary(word)
		},
	}
	return cmd
}

type DictionaryResponse struct {
	Definition string `json:"definition"`
	Word       string `json:"word"`
	Valid      bool   `json:"valid"`
}

func searchDictionary(word string) {
	fmt.Printf("meanings of [%s]\n", word)
	cfg := config.GetConfiguration()
	key := cfg.GetAPIKey(config.API_NINJAS)

	url := DICTIONARY_URL + word
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("X-Api-Key", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("server returned non-200 status: %d %s\n", resp.StatusCode, resp.Status)
		return
	}

	decoder := json.NewDecoder(resp.Body)
	var res DictionaryResponse

	if err = decoder.Decode(&res); err != nil {
		fmt.Println(err.Error())
		return
	}
	if !res.Valid {
		fmt.Println("invalid word")
		return
	}
	fmt.Println(res.Definition)
}
