package purshase

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"fmt"

	"github.com/joho/godotenv"
)

func GetBitcoinPrice() (value float64) {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error getting env. %v", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("COINMARKETCAP_API_URL")+"/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		return
	}

	q := url.Values{}
	q.Add("id", "1")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("COINMARKETCAP_API_SECRET"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to server")
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error")
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var v interface{}
	json.Unmarshal([]byte(respBody), &v)

	response := v.(map[string]interface{})
	data := response["data"]

	response = data.(map[string]interface{})
	id := response["1"]

	response = id.(map[string]interface{})
	quote := response["quote"]

	response = quote.(map[string]interface{})
	usd := response["USD"]

	response = usd.(map[string]interface{})
	price := response["price"].(float64)

	return price
}

func AmountToPrice(amount float64) float64 {
	price := GetBitcoinPrice()

	return amount * price
}
