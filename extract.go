package bca

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/shopspring/decimal"
)

// Balance is data in balance inquiry
type Balance struct {
	AccountNumber string          `json:"accountNumber"`
	Currency      string          `json:"currency"`
	BalanceRaw    string          `json:"balanceRaw"`
	Balance       decimal.Decimal `json:"balance"`
}

// Entry is a row in statement view
type Entry struct {
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
	Payee       string          `json:"payee"`
	Type        string          `json:"type"`
	AmountRaw   string          `json:"balanceRaw"`
	Amount      decimal.Decimal `json:"balance"`
}

var (
	blncCheck = regexp.MustCompile("INFORMASI REKENING - INFORMASI SALDO")
	stmtCheck = regexp.MustCompile("INFORMASI REKENING - MUTASI REKENING")
	commas    = regexp.MustCompile("[,]+")
)

const (
	blncSlctr = "#pagebody > span > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody > tr:nth-child(2)"
	stmtSlctr = "#pagebody > span > table:nth-child(2) > tbody > tr:nth-child(2) > td:nth-child(2) > table > tbody"
)

func extractHTML(b []byte, v interface{}) (err error) {
	reader := ioutil.NopCloser(bytes.NewBuffer(b))
	defer reader.Close()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	if blncCheck.Match(b) {
		return store(extractBalanceInquiry(doc), v)
	}

	if stmtCheck.Match(b) {
		return store(extractStatement(doc), v)
	}

	return fmt.Errorf("failed to extract HTML: %s %s", string(b[:]), b[:])
}

func extractBalanceInquiry(doc *goquery.Document) *Balance {
	selector := doc.Find(blncSlctr).Children()

	b := Balance{
		AccountNumber: selector.First().Text(),
		Currency:      selector.Next().First().Text(),
		BalanceRaw:    selector.Next().Next().First().Text(),
	}

	b.Balance = convDec(b.BalanceRaw)
	return &b
}

func extractStatement(doc *goquery.Document) []Entry {
	selector := doc.Find(stmtSlctr).Children().Next()

	entries := make([]Entry, 0)
	selector.Each(func(i int, s *goquery.Selection) {
		s = s.Children()
		s = s.First()

		timeRaw := s.Text()

		s = s.Next()
		desc := s.Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
			return !s.Is("br")
		})
		amountRaw := desc.Get(desc.Size() - 1).Data
		entries = append(entries, Entry{
			Date:        convTime(timeRaw),
			Description: desc.Slice(0, desc.Size()-3).Text(),
			Payee:       desc.Get(desc.Size() - 3).Data,
			Type:        s.Next().Text(),
			AmountRaw:   amountRaw,
			Amount:      convDec(amountRaw),
		})
	})
	return entries
}

func convDec(raw string) (d decimal.Decimal) {
	d, err := decimal.NewFromString(commas.ReplaceAllString(raw, ""))
	if err != nil {
		return
	}
	return d
}

func convTime(raw string) (t time.Time) {
	// Use +0700 WIB
	l, _ := time.LoadLocation("Asia/Jakarta")
	t, err := time.ParseInLocation("02/01", raw, l)
	if err != nil {
		return
	}
	return t.AddDate(time.Now().Year(), 0, 0)
}

func store(data interface{}, v interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
