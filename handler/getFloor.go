package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Payload struct {
	OperationName string    `json:"operationName"`
	Variables     Variables `json:"variables"`
	Query         string    `json:"query"`
}
type Variables struct {
	CollectionID string `json:"collectionId"`
}

func (c *Config) GetFloor() map[string]float64 {
	var floor = make(map[string]float64)
	for _, v := range c.Projects {
		id := strings.Split(v.Url, "collection/")[1]
		if strings.Contains(id, "?") {
			id = strings.Split(id, "?")[0]
		}
		data := []byte(`{"operationName":"GetCollection","variables":{"collectionId":"` + id + `"},"query":"query GetCollection($collectionId: ID!) {\n  public {\n    collection(id: $collectionId) {\n      id\n      name\n      description\n      categories\n      banner {\n        url\n        __typename\n      }\n      logo {\n        url\n        __typename\n      }\n      creator {\n        displayName\n        __typename\n      }\n      aggregatedAttributes {\n        label: traitType\n        options: attributes {\n          value: id\n          label: value\n          total\n          __typename\n        }\n        __typename\n      }\n      metrics {\n        items\n        minAuctionListingPriceDecimal\n        minSaleListingPriceDecimal\n        owners\n        totalSalesDecimal\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"}`)

		req, err := http.NewRequest("POST", "https://crypto.com/nft-api/graphql", bytes.NewBuffer(data))
		if err != nil {
			// handle err
		}

		req.Header.Set("Authority", "crypto.com")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Apollographql-Client-Name", "web")
		req.Header.Set("Apollographql-Client-Version", "current")
		req.Header.Set("B3", "f98e8de006d5b79904d06b51bdd868c2-3c93f866db21ec88-1")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Cookie", "__cf_bm=.F_h.EQqgncwCSVRwYH2r8Nl9MPhlX2JPnXt5kc1iZU-1652635948-0-AQ5oEq3ONy+hAnbtSCwh5AH+EtuTUbpszVDJ7klso5XL2t2LnB7+tstS7qkqF2TK8F3OEg1G+3B7BFvuG+0wx2Q=; _gcl_au=1.1.319662932.1652635949; __55=%7B%22ms%22%3A%22non-member%22%2C%22st%22%3A%22regular%22%2C%22vF0%22%3A1652635948623%2C%22vF%22%3A%22new%22%7D; _rdt_uuid=1652635948721.431b1757-0c1c-48a6-982f-d52f38c72a88; _gid=GA1.2.836800589.1652635949; _scid=e34d7eea-bd50-4164-862b-325f48b7245b; _tt_enable_cookie=1; _ttp=5aeeda13-8417-43be-bf53-95225f97e960; _fbp=fb.1.1652635948931.865645072; _pin_unauth=dWlkPVpESTFOVGcwT0RZdE56ZzNaaTAwWVdRNUxUaG1ZalF0TURjeFpUazNZamszTkRreQ; _sctr=1|1652587200000; OptanonConsent=isIABGlobal=false&datestamp=Sun+May+15+2022+13%3A32%3A29+GMT-0400+(Eastern+Daylight+Time)&version=6.5.0&hosts=&landingPath=https%3A%2F%2Fcrypto.com%2Fnft%2Fcollection%2F4ff90f089ac3ef9ce342186adc48a30d&groups=C0001%3A1%2CC0002%3A0%2CC0003%3A0%2CC0004%3A0; __ssid=f6bf08a1b833619257a02f8443355ec; intercom-id-ruozuwky=5c62f30f-d563-4eb1-be97-ea8c0d4549d2; intercom-session-ruozuwky=; _ga_KTR8M2WC2H=GS1.1.1652635948.1.1.1652636076.60; _gat__ga=1; _ga_PC0RNJG7RR=GS1.1.1652635948.1.1.1652636076.60; _ga=GA1.1.909660251.1652635949")
		req.Header.Set("Origin", "https://crypto.com")
		req.Header.Set("Referer", "https://crypto.com/nft/collection/4ff90f089ac3ef9ce342186adc48a30d?buyNow=true&sort=price&order=ASC")
		req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
		req.Header.Set("X-Queueit-Ajaxpageurl", "https%3A%2F%2Fcrypto.com%2Fnft%2Fcollection%2F4ff90f089ac3ef9ce342186adc48a30d%3FbuyNow%3Dtrue%26sort%3Dprice%26order%3DASC")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
		}
		if resp.StatusCode != 200 {
			fmt.Println(resp.Status)
			return nil
		}
		defer resp.Body.Close()
		var floorprice FloorData
		err = json.NewDecoder(resp.Body).Decode(&floorprice)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		floor[v.Name] = StringToFloat(floorprice.Data.Public.Collection.Metrics.MinSaleListingPriceDecimal)
	}
	return floor

}
func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return f
}
