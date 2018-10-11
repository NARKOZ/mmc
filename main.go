package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type CurrencyData struct {
	Currencies map[string]Currency `json:"results"`
}

type Currency struct {
	Name string `json:"currencyName"`
}

func loadCurrencies() CurrencyData {
	var data CurrencyData

	// List of all currencies from http://free.currencyconverterapi.com/api/v3/currencies
	file, _ := Asset("data/currencies.json")
	json.Unmarshal(file, &data)

	return data
}

func getCurrencyNames(from string, to string) (string, string) {
	data := loadCurrencies()
	fromCurrencyName := data.Currencies[from].Name
	toCurrencyName := data.Currencies[to].Name

	return fromCurrencyName, toCurrencyName
}

func isValidCurrency(currencyID string) bool {
	data := loadCurrencies()
	currencyName := data.Currencies[currencyID].Name

	return currencyName != ""
}

func handleError(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func getRate(rateID string) float64 {
	url := "http://free.currencyconverterapi.com/api/v3/convert?q=" + rateID + "&compact=ultra"

	response, err := http.Get(url)
	if err != nil {
		handleError("Error getting data")
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		handleError("Error fetching data")
	}

	rate := data[rateID]
	if rate == nil {
		handleError("No results for currency rate " + rateID)
	}

	return rate.(float64)
}

func listSupportedCurrencies() {
	fmt.Println("Supported currencies:")
	for id, currency := range loadCurrencies().Currencies {
		fmt.Printf("%s\t%s\n", id, currency.Name)
	}
}

func parseArgs(args []string) (float64, string, string) {

	usageList := "Use \"mmc list\" to list all currencies supported by mmc."
	usage := "Usage: mmc <amount> <source_currency> to <target_currency>\n" +
		"Example: mmc 100 USD to AUD\n" + usageList

	if len(args) == 2 && args[1] == "list" {
		listSupportedCurrencies()
		os.Exit(0)
	}

	if len(args) < 5 {
		handleError(fmt.Sprintf("Insufficient arguments\n%s", usage))
	}

	from, to := strings.ToUpper(args[2]), strings.ToUpper(args[4])
	if !isValidCurrency(from) {
		handleError(fmt.Sprintf("Invalid or unsupported currency: %s\n\n%s", from, usageList))
	}
	if !isValidCurrency(to) {
		handleError(fmt.Sprintf("Invalid or unsupported currency: %s\n\n%s", to, usageList))
	}

	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		handleError(fmt.Sprintf("Invalid value for conversion: %s", args[1]))
	}

	return amount, from, to
}

func main() {

	amount, from, to := parseArgs(os.Args)

	result := amount * getRate(from+"_"+to)
	fromCurrency, toCurrency := getCurrencyNames(from, to)

	fmt.Println(amount, fromCurrency, "=", result, toCurrency)
}
