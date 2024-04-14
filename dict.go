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
  Definition string `json:"definition"`
}

func getMeaning(encodedURL string) *http.Response  {
  response, error := http.Get(encodedURL)
  if error != nil {
    log.Fatalf("Error calling the url %v ", error)
  }
  return response
}

func parseAndPrint(dictionary []Dictionary)  {
  fmt.Printf("Word -> %s \n", dictionary[0].Word)
  fmt.Printf("Phonetic -> %s \n", dictionary[0].Phonetic)
  definitions := dictionary[0].Meanings[0].Definitions

  fmt.Println("Definitions ")
  for _, definition := range definitions {
    fmt.Println("\t| ")
    fmt.Printf("\t -> %s \n", definition.Definition)
  }
}

func main()  {

  words := os.Args[1:]
  if len(words) == 0 {
    log.Fatalf("Please enter a word get the meaning")
  }

  encodedURL := fmt.Sprintf(url, words[0])
  response := getMeaning(encodedURL)
  defer response.Body.Close()

  body, error := io.ReadAll(response.Body)
  if error != nil {
    fmt.Printf("Error reading the meaning from response -> %s ", error)
    return
  }

  var dictionary []Dictionary
  error = json.Unmarshal(body, &dictionary)
  if len(dictionary) == 0 {
    log.Fatalf("Unable to read meaning for the word -> %s ", words[0])
  }

  parseAndPrint(dictionary)

}
