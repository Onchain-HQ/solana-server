package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NativeBalanceRPCResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			Slot int64 `json:"slot"`
		} `json:"context"`
		Value int64 `json:"value"` // This holds the balance
	} `json:"result"`
	Id int64 `json:"id"`
}

func GetSolBalance(address string) (string, error) {
	apiUrl := "https://api.mainnet-beta.solana.com/"

	jsonData := []byte(fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "getBalance",
		"params": ["%s"]
	}`, address))

	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "0", fmt.Errorf("Failed to fetch native balance: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	if resp.StatusCode != http.StatusOK {
		// Read the response body for error details
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		return "0", fmt.Errorf("Failed to fetch native balance, status code: %d, response: %s\n", resp.StatusCode, bodyString)
	}

	var res NativeBalanceRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "0", err
	}

	return fmt.Sprintf("%d", res.Result.Value), nil

}

type ExchangeRateResponse struct {
	Data struct {
		Currency string            `json:"currency"`
		Rates    map[string]string `json:"rates"`
	} `json:"data"`
}

func GetExchangeRateInfo() (*ExchangeRateResponse, error) {
	apiUrl := "https://api.coinbase.com/v2/exchange-rates?currency=USD"

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate data, status code: %v", err.Error())
	}
	defer resp.Body.Close()

	var coinbaseRes *ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&coinbaseRes); err != nil {
		return nil, fmt.Errorf("Failed to decode exchange rate data: %v", err)
	}
	return coinbaseRes, nil
}
