package rocketfuel

import (
	"fmt"
	"testing"
	"time"
)

var client *Client

func init() {

	// options := &Options{
	// 	Environment: "qa",
	// 	PublicKey: `-----BEGIN PUBLIC KEY-----
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxyJvURWxnf96J8/5JdSB
	// clfC8ZEhbU48hYZYsm5J9f2HXsX6Ttoo6KEfB0XKPg674FEDVlDtG8Cq9m2dsDzx
	// Ekp05dj6TiPgxc9b2dmom0iff6rq05gWumF+CDvQW9euawR1gIutWH36EgKbyW1j
	// 3lvYAihfGj5FOk9DDkaQPYtyuID9G69ljwzFhR+Zbpjyf5GWd+3vEipLkAo8CbXs
	// l2NPrAlil3a9xl0OnqMLGncGCIKxJnvGEH7MakAepD8rDl3EKFiLuMEv65y2hLvj
	// KAADkD8kP1xTiAwYYLLDz2TFEkqO2iM6XmZkugD1u7F6zJ+1A5J+WxBy/NUmLU6A
	// JQIDAQAB
	// -----END PUBLIC KEY-----`,
	// 	ClientId:     "6e79b45cff6aaf50289a78b85648018f0e6590be798bf03d0dc0cb04440fef7d",
	// 	MerchantId:   "4384ee2f-79df-4f25-90f1-700147b4adc4",
	// 	ClientSecret: "23956ed6-33b0-452f-98f7-6b558db062b6",
	// }
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
