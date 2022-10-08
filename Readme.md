# Go library for the Rocketfuel API.

rkfl-sdk-go is a Go client library for accessing the Rocketfuel API.

Where possible, the services available on the client groups the API into logical chunks and correspond to the structure of the Rocketfuel API documentation at https://docs.rocketfuelblockchain.com/developer-guides/api-reference

## Usage

``` go
import "https://bitbucket.org/rocketfuelblockchain/rocketfuel-sdk-go"

options := &rocketfuel.Options{
		Environment: "sandbox",
		PublicKey: "MERCHANT_PUBLICKEY",
		Email:      "MERCHANT_EMAIL",
		MerchantId: "MERCHANT_ID",
		Password:   "MERCHANT_PASSWORD",
	}

// second param is an optional http client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
client := rocketfuel.NewClient(options, nil)

cart1 := rocketfuel.Cart{
    Id:       "Test",
    Name:     "Test",
    Price:    "10",
    Quantity: "2",
}
payload := rocketfuel.HostedPageRequest{
    Amount:      "20",
    Cart:        []rocketfuel.Cart{cart1},
    Currency:    "USD",
    Order:       string("ORDER_ID"),
    Merchant_id: "MERCHANT_ID",
    RedirectUrl: "",
}
result, _ := client.GetUUID(payload)

fmt.Println("Result:", result)
```

## TODO
- [ ] Use option-supplied merchant Id for all requests
- [ ] Tests
