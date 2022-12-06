package rocketfuel

import (
	"fmt"
	"testing"
	"time"
)

var client *Client

func init() {
	options := &Options{
		Environment:  "sandbox",
		PublicKey:    `-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAslC4WPywoZKINSmL4aMs\nT94/4LOinE75SuY3TB6BNxnqylxTydu6/X2HCHofdw1VDzm1fi2/9FtWZHzGRro5\nTzbdoHqT/Nc0hIkm8MuiU2FjT+E9G/JAGZGXdHUVp9ti29HzOTdYFNYppKr5XYJQ\n3qseKPpUUwQR1dhYC4mDWSp3JTjKz4opN6zRSeTw1Dr6PtnqgfRzkTDQbMK4uWvp\nbJ0/SFuwpBr+tgHxMhaCImdQd3Zck0M7CcS1bBCQDbSYVU/gS5wig76lycHCK0Xq\nS5l38X+geJ9wR9oiUEVluh9KASGU+jc3SEWo0+WJM03d2OXk+Y7ObtCi5D35Y/e3\n3QIDAQAB\n-----END PUBLIC KEY-----`,
		ClientId:     "45f880085d700cab1b16a506357b6bc4459b49864933ce6a91a47f2863f630c7",
		MerchantId:   "9514ec97-8672-4668-bf43-8722c9fe89c2",
		ClientSecret: "fdb652e4-ee6e-478d-b332-53f9c045663b",
	}

	client = NewClient(options, nil)
}
func TestUUID(t *testing.T) {

	cart1 := Cart{
		Id:       "Test",
		Name:     "Test",
		Price:    "10",
		Quantity: "2",
	}
	payload := HostedPageRequest{
		Amount:      "20",
		Cart:        []Cart{cart1},
		Currency:    "USD",
		Order:       time.Now().String(),
		RedirectUrl: "",
	}

	result, _ := client.GetUUID(payload)

	fmt.Println("Cart Result:", result)
}
