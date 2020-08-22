# bca-go

go-bca is the (unofficial) BCA SDK for the Go programming language.

This SDK was kickstarted by the [OpenAPI Generator](https://openapi-generator.tech) project.

## Installation

```shell
go get github.com/satraul/bca-go
```

## Usage

```golang
import (
    bca "github.com/satraul/bca-go"
)

func main() {
	api := bca.NewAPIClient(bca.NewConfiguration())
    auth, err := api.Login(ctx, "username", "password", "1.2.3.4")
	if err != nil {
		panic(err)
	}

	balance, err := api.BalanceInquiry(ctx, auth)
	if err != nil {
		panic(err)
    }
    log.Printf("%+v\n", balance)

    if err := api.Logout(ctx, auth); err != nil {
        panic(err)
    }
}
```

See the a full example at [example/example.go](example/example.go).

## Documentation for API Endpoints

All URIs are relative to *https://m.klikbca.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BCAApi* | [**AccountStatementView**](docs/BCAApi.md#accountstatementview) | **Post** /accountstmt.do?value(actions)&#x3D;acctstmtview | AccountStatementView
*BCAApi* | [**BalanceInquiry**](docs/BCAApi.md#balanceinquiry) | **Post** /balanceinquiry.do | BalanceInquiry
*BCAApi* | [**Login**](docs/BCAApi.md#login) | **Post** /authentication.do | Login
*BCAApi* | [**Logout**](docs/BCAApi.md#logout) | **Get** /authentication.do?value(actions)&#x3D;logout | Logout

## Documentation For Authorization

[**Login**](docs/BCAApi.md#login) will return session cookies ([]*Cookie) that are used for auth in all other endpoints.

## Contributing

Pull requests are welcome.

## License

[MIT](LICENSE)