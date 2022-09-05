package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ReqBody struct {
	SourceLang string `json:"sourceLang"`
	TargetLang string `json:"targetLang"`
	SourceText string `json:"sourceText"`
}

type Resp struct {
	Data struct {
		Languages []struct {
			Language string `json:"language"`
		} `json:"languages"`
	} `json:"data"`
}

const translateURL = "https://google-translate1.p.rapidapi.com/language/translate/v2"

func ReqTranslate(body *ReqBody) ([]byte, error) {
	var str2 string
	str2 = ""
	str2 = str2 + "q=" + body.SourceText
	str2 = str2 + "&target=" + body.TargetLang
	str2 = str2 + "&source=" + body.SourceLang

	payload := strings.NewReader(str2)

	req, err := http.NewRequest("POST", translateURL, payload)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "1f567f5ad9msh566e3bdad55c972p12adadjsnbca1e41ae0e7")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer res.Body.Close()
	body1, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		return []byte(""), errors.New("Too many requests")
	}
	return body1, nil
}

func GetLanguages() ([]string, error) {
	var language []string
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2/languages"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return language, err
	}
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "1f567f5ad9msh566e3bdad55c972p12adadjsnbca1e41ae0e7")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error in request", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return language, err
	}
	var rs Resp
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return language, err
	}
	for _, v := range rs.Data.Languages {
		language = append(language, v.Language)
	}
	return language, nil
}
