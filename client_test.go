package rocketfuel

import (
	"fmt"
	"testing"
	"time"
)
var client *Client
func init(){

	options := &Options{
		Environment: "qa",
		PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxyJvURWxnf96J8/5JdSB
clfC8ZEhbU48hYZYsm5J9f2HXsX6Ttoo6KEfB0XKPg674FEDVlDtG8Cq9m2dsDzx
Ekp05dj6TiPgxc9b2dmom0iff6rq05gWumF+CDvQW9euawR1gIutWH36EgKbyW1j
3lvYAihfGj5FOk9DDkaQPYtyuID9G69ljwzFhR+Zbpjyf5GWd+3vEipLkAo8CbXs
l2NPrAlil3a9xl0OnqMLGncGCIKxJnvGEH7MakAepD8rDl3EKFiLuMEv65y2hLvj
KAADkD8kP1xTiAwYYLLDz2TFEkqO2iM6XmZkugD1u7F6zJ+1A5J+WxBy/NUmLU6A
JQIDAQAB
-----END PUBLIC KEY-----`,
		ClientId:      "6e79b45cff6aaf50289a78b85648018f0e6590be798bf03d0dc0cb04440fef7d",
		MerchantId: "4384ee2f-79df-4f25-90f1-700147b4adc4",
		ClientSecret:   "23956ed6-33b0-452f-98f7-6b558db062b6",
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
		Amount:   "20",
		Cart:     []Cart{cart1},
		Currency: "USD",
		Order:    time.Now().String(),
		// Merchant_id: "4384ee2f-79df-4f25-90f1-700147b4adc4",
		RedirectUrl: "",
	}
 
	result, _ := client.GetUUID(payload)

	fmt.Println("Cart Result:", result)
}
