package main

import (
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestParseArgs(t *testing.T) {
	args := []string{"", "9000.3", "eur", "to", "usd"}
	amount, from, to := parseArgs(args)

	if amount != 9000.3 {
		t.Fatalf("Expected '9000.3' got '%v'", amount)
	}

	if from != "EUR" {
		t.Fatalf("Expected 'EUR' got '%v'", from)
	}

	if to != "USD" {
		t.Fatalf("Expected 'USD' got '%v'", to)
	}
}

func TestLoadCurrencies(t *testing.T) {
	data := loadCurrencies()

	if data.Currencies == nil {
		t.Fatalf("Expected 'Currencies' struct got nil")
	}
}

func TestGetCurrencyNames(t *testing.T) {
	name1, name2 := getCurrencyNames("EUR", "USD")

	if name1 != "European Euro" {
		t.Fatalf("Expected 'European Euro' got '%v'", name1)
	}

	if name2 != "United States dollar" {
		t.Fatalf("Expected 'United States dollar' got '%v'", name2)
	}
}

func TestIsValidCurrency(t *testing.T) {
	valid := isValidCurrency("USD")
	invalid := isValidCurrency("invalid")

	if valid != true {
		t.Fatalf("Expected 'true' got '%v'", valid)
	}

	if invalid != false {
		t.Fatalf("Expected 'false' got '%v'", invalid)
	}
}

func TestGetRate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://free.currencyconverterapi.com/api/v3/convert?q=EUR_USD&compact=ultra",
		httpmock.NewStringResponder(200, `{"EUR_USD":700.007}`))

	rate := getRate("EUR_USD")

	if rate != 700.007 {
		t.Fatalf("Expected '700.007' got '%v'", rate)
	}
}
