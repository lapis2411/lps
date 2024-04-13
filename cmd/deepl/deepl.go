package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	config "github.com/lapis2411/tools/config"
)

var DeeplCmd = &cobra.Command{
	Use:   "deepl",
	Short: "translate the input sentence",
	Long:  `translate the input sentence`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		translate(strings.Join(args, " "))
	},
}

type RequestData struct {
	Text       []string `json:"text"`
	TargetLang string   `json:"target_lang"`
}

type TranslateData struct {
	SourceLanguage string `json:"detected_source_language"`
	Text           string `json:"text"`
}

type ExpectedResponseData struct {
	Translations []TranslateData `json:"translations"`
}

func translate(text string) {
	fmt.Println("text:", text)
	cfg := config.GetConfiguration()
	key := cfg.GetAPIKey(config.DEEPL)

	request := RequestData{
		Text:       []string{text},
		TargetLang: "JA",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Authorization", "DeepL-Auth-Key "+key)
	req.Header.Add("Content-Type", "application/json")

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
	var exp ExpectedResponseData

	if err = decoder.Decode(&exp); err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range exp.Translations {
		fmt.Println(v.Text)
	}
}
