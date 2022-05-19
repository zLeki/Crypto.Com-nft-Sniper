package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (c *Config) GetNfts() ([]NftData, error) {

	var nfts []NftData
	for _, v := range c.Projects {
		type Payload struct {
			OperationName string    `json:"operationName"`
			Variables     Variables `json:"variables"`
			Query         string    `json:"query"`
		}
		type Where struct {
			AssetName   interface{}   `json:"assetName"`
			Description interface{}   `json:"description"`
			MinPrice    interface{}   `json:"minPrice"`
			MaxPrice    interface{}   `json:"maxPrice"`
			BuyNow      bool          `json:"buyNow"`
			Auction     bool          `json:"auction"`
			Attributes  []interface{} `json:"attributes"`
			Chains      []interface{} `json:"chains"`
		}
		type Sort struct {
			Order string `json:"order"`
			Field string `json:"field"`
		}
		type Variables struct {
			CollectionID string `json:"collectionId"`
			First        int    `json:"first"`
			Skip         int    `json:"skip"`
			CacheID      string `json:"cacheId"`
			Where        Where  `json:"where"`
			Sort         []Sort `json:"sort"`
		}

		data := []byte(`{"operationName":"GetAssets","variables":{"collectionId":"` + strings.Split(v.Url, "collection/")[1] + `","first":6,"skip":0,"cacheId":"getAssetsQuery-0a2c38bb89f6209c0ea3578bb4a8631a1a9ecd37","where":{"assetName":null,"description":null,"minPrice":null,"maxPrice":null,"buyNow":true,"auction":false,"attributes":[],"chains":[]},"sort":[{"order":"ASC","field":"price"}]},"query":"fragment UserData on User {\n  uuid\n  id\n  username\n  displayName\n  isCreator\n  avatar {\n    url\n    __typename\n  }\n  isCreationWithdrawalBlocked\n  creationWithdrawalBlockExpiredAt\n  verified\n  __typename\n}\n\nquery GetAssets($audience: Audience, $brandId: ID, $categories: [ID!], $collectionId: ID, $creatorId: ID, $ownerId: ID, $first: Int!, $skip: Int!, $cacheId: ID, $hasSecondaryListing: Boolean, $where: AssetsSearch, $sort: [SingleFieldSort!], $isCurated: Boolean, $createdPublicView: Boolean) {\n  public(cacheId: $cacheId) {\n    assets(\n      audience: $audience\n      brandId: $brandId\n      categories: $categories\n      collectionId: $collectionId\n      creatorId: $creatorId\n      ownerId: $ownerId\n      first: $first\n      skip: $skip\n      hasSecondaryListing: $hasSecondaryListing\n      where: $where\n      sort: $sort\n      isCurated: $isCurated\n      createdPublicView: $createdPublicView\n    ) {\n      id\n      name\n      copies\n      copiesInCirculation\n      creator {\n        ...UserData\n        __typename\n      }\n      main {\n        url\n        __typename\n      }\n      cover {\n        url\n        __typename\n      }\n      royaltiesRateDecimal\n      primaryListingsCount\n      secondaryListingsCount\n      primarySalesCount\n      totalSalesDecimal\n      defaultListing {\n        editionId\n        priceDecimal\n        mode\n        auctionHasBids\n        __typename\n      }\n      defaultAuctionListing {\n        editionId\n        priceDecimal\n        auctionMinPriceDecimal\n        auctionCloseAt\n        mode\n        auctionHasBids\n        __typename\n      }\n      defaultSaleListing {\n        editionId\n        priceDecimal\n        mode\n        __typename\n      }\n      defaultPrimaryListing {\n        editionId\n        priceDecimal\n        mode\n        auctionHasBids\n        primary\n        __typename\n      }\n      defaultSecondaryListing {\n        editionId\n        priceDecimal\n        mode\n        auctionHasBids\n        __typename\n      }\n      defaultSecondaryAuctionListing {\n        editionId\n        priceDecimal\n        auctionMinPriceDecimal\n        auctionCloseAt\n        mode\n        auctionHasBids\n        __typename\n      }\n      defaultSecondarySaleListing {\n        editionId\n        priceDecimal\n        mode\n        __typename\n      }\n      likes\n      views\n      isCurated\n      defaultEditionId\n      isLiked\n      isAcceptingOffers\n      isExternalNft\n      externalNftMetadata {\n        creatorAddress\n        creator {\n          name\n          avatar {\n            url\n            __typename\n          }\n          __typename\n        }\n        network\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"}`)

		req, err := http.NewRequest("POST", "https://crypto.com/nft-api/graphql", bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authority", "crypto.com")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Apollographql-Client-Name", "web")
		req.Header.Set("Apollographql-Client-Version", "current")
		req.Header.Set("B3", "0183f0b215c3e6e6074be847ef96d400-5936d81732324591-1")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Cookie", "_gcl_au=1.1.319662932.1652635949; _rdt_uuid=1652635948721.431b1757-0c1c-48a6-982f-d52f38c72a88; _scid=e34d7eea-bd50-4164-862b-325f48b7245b; _tt_enable_cookie=1; _ttp=5aeeda13-8417-43be-bf53-95225f97e960; _fbp=fb.1.1652635948931.865645072; _pin_unauth=dWlkPVpESTFOVGcwT0RZdE56ZzNaaTAwWVdRNUxUaG1ZalF0TURjeFpUazNZamszTkRreQ; _sctr=1|1652587200000; __ssid=f6bf08a1b833619257a02f8443355ec; intercom-id-ruozuwky=5c62f30f-d563-4eb1-be97-ea8c0d4549d2; intercom-session-ruozuwky=; __cf_bm=oswus6LEMGamdNchtwvwPpVdeM51DUZghqc4kEaJNBY-1652910683-0-AdrnC97t9HRZCg7U/UHpwmsUOuyvwe/t2EJfAu6efnKvWZj2gvmNCMVc3UwTItQp9hWc7jbgvc+XOtOxiOIo4Ns=; __55=%7B%22ms%22%3A%22non-member%22%2C%22st%22%3A%22regular%22%2C%22vF0%22%3A1652635948623%2C%22vF%22%3A%22occasional%22%2C%22vF1%22%3A1652910684350%7D; _gid=GA1.2.1731681371.1652910684; _gat__ga=1; _gcl_aw=GCL.1652910695.Cj0KCQjwspKUBhCvARIsAB2IYus5mEaR034qygbtkvoeJ5jLg5Znw5WKfvx9Z9Qdjb4SbQvUIzSeW90aApfREALw_wcB; _gac_UA-99317940-18=1.1652910695.Cj0KCQjwspKUBhCvARIsAB2IYus5mEaR034qygbtkvoeJ5jLg5Znw5WKfvx9Z9Qdjb4SbQvUIzSeW90aApfREALw_wcB; _ga=GA1.1.909660251.1652635949; OptanonAlertBoxClosed=2022-05-18T21:51:45.270Z; OptanonConsent=isIABGlobal=false&datestamp=Wed+May+18+2022+17%3A51%3A45+GMT-0400+(Eastern+Daylight+Time)&version=6.5.0&hosts=&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1%2CC0004%3A1&AwaitingReconsent=false; _ga_KTR8M2WC2H=GS1.1.1652910684.2.1.1652910717.27; _ga_PC0RNJG7RR=GS1.1.1652910684.2.1.1652910717.27")
		req.Header.Set("Origin", "https://crypto.com")
		req.Header.Set("Referer", "https://crypto.com/nft/collection/82421cf8e15df0edcaa200af752a344f?buyNow=true&sort=price&order=ASC")
		req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
		req.Header.Set("X-Queueit-Ajaxpageurl", "https%3A%2F%2Fcrypto.com%2Fnft%2Fcollection%2F82421cf8e15df0edcaa200af752a344f%3FbuyNow%3Dtrue%26sort%3Dprice%26order%3DASC")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("Not 200")
		}
		var singleNft NftData
		err = json.NewDecoder(resp.Body).Decode(&singleNft)
		if err != nil {
			return nil, err
		}

		nfts = append(nfts, singleNft)
	}
	return nfts, nil
}
