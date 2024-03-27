package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "io"
  "encoding/json"
)

const url string = "https://api.dictionaryapi.dev/api/v2/entries/en/%s"
type Dictionary struct {
  Word string `json:"word"`
  Phonetic string `json:"phonetic"`
  Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
  Definitions []Definition `json:"definitions"`
}

type Definition struct {
  Definition string `json:"definition`
}

func getMeaning(encodedURL string) *http.Response  {
  response, error := http.Get(encodedURL)
  if error != nil {
    log.Fatalf("Error calling the url %v", error)
  }
  return response
}

func parseAndPrint(dictionary []Dictionary)  {
  for _, word := range dictionary {
    fmt.Println("Word: ", word.Word)
    fmt.Println("Phonetic: ", word.Phonetic)
    fmt.Println("Meaning: ", word.Meanings[0].Definitions[0].Definition)
  }
}

func main()  {

  words := os.Args[1:]
  encodedURL := fmt.Sprintf(url, words[0])

  response := getMeaning(encodedURL)
  defer response.Body.Close()

  body, error := io.ReadAll(response.Body)
  if error != nil {
    fmt.Printf("Error reading the meaning from response %s", error)
    return
  }

  var dictionary []Dictionary
  error = json.Unmarshal(body, &dictionary)
  if error != nil {
    log.Fatalf(" word %v", error)
  }

  parseAndPrint(dictionary)

}
