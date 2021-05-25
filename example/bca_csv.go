package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	bca "github.com/satraul/bca-go"
)

const (
	TypePositive string = "CR"
	TypeNegative        = "DB"
)

type TransactionRow struct {
	Amount     string `csv:"amount"`
	IsPositive bool   `csv:"is_positive"`
	Memo       string `csv:"memo"`
	Date       string `csv:"date"`
	UniqueID   string `csv:"unique_id"`
}

func main() {
	var (
		ctx      = context.Background()
		username = os.Getenv("BCA_USERNAME")
		password = os.Getenv("BCA_PASSWORD")
		ip       = getPublicIP()
	)
	api := bca.NewAPIClient(bca.NewConfiguration())
	auth, err := api.Login(ctx, username, password, ip)
	if err != nil {
		panic(err)
	}
	defer api.Logout(ctx, auth)

	entries, err := api.AccountStatementView(ctx, time.Now().AddDate(0, 0, -7), time.Now().AddDate(0, 0, -6), auth)
	if err != nil {
		log.Fatal(err)
		return
	}

	var rows []*TransactionRow
	for _, entry := range entries {
		rows = append(rows, &TransactionRow{
			Amount:     strings.Replace(entry.Amount.StringFixed(2), ".", ",", -1),
			IsPositive: entry.Type == TypePositive,
			Memo:       entry.Description,
			Date:       entry.Date.Format("02/1/2006"),
		})
	}

	// save to csv
	now := time.Now()
	fileName := fmt.Sprintf("%s.csv", now.Format("2006-01-02T15.04.05"))
	csvFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	err = gocsv.MarshalFile(&rows, csvFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}

// https://gist.github.com/ankanch/8c8ec5aaf374039504946e7e2b2cdf7f
func getPublicIP() string {
	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(ip)
}
