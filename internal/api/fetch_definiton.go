package api

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "os"
)

type APIResponse []struct {
	Hwi struct {
		Prs []struct {
			Mw string `json:"mw"`
		} `json:"prs,omitempty"`
	} `json:"hwi"`
	Fl       string   `json:"fl"`       
	Shortdef []string `json:"shortdef"` 
}


type DefinitonResult struct {
    Pronunciations  []string
    WordType 		string
    Definitions     []string
}


const websterbaseURL = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"

func FetchDefinition(ctx context.Context, word string) (*DefinitonResult, error) {
    apiKey := os.Getenv("API_KEY")

    if apiKey == "" {
        return nil, fmt.Errorf("API_KEY is not set. Please set the API key in Env Variable to proceed using the API")
    }

	url := fmt.Sprintf("%s%s?key=%s", websterbaseURL, word, apiKey)
	

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching data:", err)
	}

	defer resp.Body.Close()


    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned non-200: %s", resp.Status)
    }		


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response:", err)
	}

	var responseObject  APIResponse

	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		return nil,fmt.Errorf("The supplied word '%s' doesn't exists in the dictionary. Please check and try again",word)
	}

    if len(responseObject) == 0 || len(responseObject[0].Shortdef) == 0 {
        return nil, fmt.Errorf("no definitions found for '%s'", word)
    }


	var pronunciations []string
    for _, p := range responseObject[0].Hwi.Prs {
        if p.Mw != "" {
            pronunciations = append(pronunciations, p.Mw)
        }
    }

    return &DefinitonResult{
        Pronunciations:  pronunciations,
        WordType: responseObject[0].Fl,
        Definitions:     responseObject[0].Shortdef,
    }, nil	



}


	




	
	


	