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
	ctx := context.Background()
	bcaclient := bca.NewAPIClient(bca.NewConfiguration())
	var (
		username = os.Getenv("BCA_USERNAME")
		password = os.Getenv("BCA_PASSWORD")
		ip       = getPublicIP()
	)

	_, r, err := bcaclient.BCAApi.Login(ctx, username, password, ip)
	if err != nil {
		log.Fatal(err)
	}
	auth := r.Cookies()

	balance, _, err := bcaclient.BCAApi.BalanceInquiry(ctx, auth)
	if err != nil {
		bcaclient.BCAApi.Logout(ctx, auth)
		log.Fatal(err)
	}
	log.Printf("%+v\n", balance)

	statement, _, err := bcaclient.BCAApi.AccountStatementView(ctx, time.Now().AddDate(0, 0, -7), time.Now(), auth)
	if err != nil {
		bcaclient.BCAApi.Logout(ctx, auth)
		log.Fatal(err)
	}
	log.Printf("%+v\n", statement)

	bcaclient.BCAApi.Logout(ctx, auth)
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
