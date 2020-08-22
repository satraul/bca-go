package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	bca "github.com/satraul/bca-go"
)

func main() {
	retcode := 1
	defer func() { os.Exit(retcode) }()

	var (
		ctx      = context.Background()
		username = os.Getenv("BCA_USERNAME")
		password = os.Getenv("BCA_PASSWORD")
		ip       = getPublicIP()
	)

	api := bca.NewAPIClient(bca.NewConfiguration())
	auth, err := api.Login(ctx, username, password, ip)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer api.Logout(ctx, auth)

	balance, err := api.BalanceInquiry(ctx, auth)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%+v\n", balance)

	statement, err := api.AccountStatementView(ctx, time.Now().AddDate(0, 0, -7), time.Now(), auth)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%+v\n", statement)

	retcode = 0
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
