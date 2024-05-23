package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type commonResponse struct {
	Error  string   `json:"error,omitempty"`
	Result []string `json:"result,omitempty"`
	Type   string   `json:"type,omitempty"`
}

func process_show_dashboard(dashboardName string) (result []byte, cType string, err error) {

	url := "http://10.13.57.9:8080/api/v1/dashboard/show"

	payload := []byte(fmt.Sprintf(`{"command":"%s"}`, dashboardName))

	fmt.Println(string(payload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, "", err
	}

	mimeType := http.DetectContentType(body)

	if resp.StatusCode != 200 {
		return nil, "", errors.New("no dashboard name found")
	}

	return body, mimeType, nil
}

func process_list_dashboard() (listDashboard []string, err error) {

	url := "http://localhost:8080/api/v1/dashboard/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var commonResponse commonResponse
	err = json.Unmarshal(body, &commonResponse)
	if err != nil {
		return nil, err
	}

	return commonResponse.Result, nil
}

func process_list_dashboard_by_cat(category string) (listDashboard []string, err error) {

	urlCategory := "http://localhost:8080/api/v1/dashboard/category"

	urlPath, err := url.Parse(urlCategory)
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Add("category", category)
	urlPath.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", urlPath.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var commonResponse commonResponse
	err = json.Unmarshal(body, &commonResponse)
	if err != nil {
		return nil, err
	}

	return commonResponse.Result, nil
}
