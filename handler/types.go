package handler

type Config struct {
	Projects []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"projects"`
	Solanawalletkey string `json:"solanawalletkey"`
}
type FloorData struct {
	Data struct {
		Public struct {
			Collection struct {
				ID          string   `json:"id"`
				Name        string   `json:"name"`
				Description string   `json:"description"`
				Categories  []string `json:"categories"`
				Banner      struct {
					URL      string `json:"url"`
					Typename string `json:"__typename"`
				} `json:"banner"`
				Logo struct {
					URL      string `json:"url"`
					Typename string `json:"__typename"`
				} `json:"logo"`
				Creator struct {
					DisplayName string `json:"displayName"`
					Typename    string `json:"__typename"`
				} `json:"creator"`
				AggregatedAttributes []struct {
					Label   string `json:"label"`
					Options []struct {
						Value    string `json:"value"`
						Label    string `json:"label"`
						Total    int    `json:"total"`
						Typename string `json:"__typename"`
					} `json:"options"`
					Typename string `json:"__typename"`
				} `json:"aggregatedAttributes"`
				Metrics struct {
					Items                         int    `json:"items"`
					MinAuctionListingPriceDecimal string `json:"minAuctionListingPriceDecimal"`
					MinSaleListingPriceDecimal    string `json:"minSaleListingPriceDecimal"`
					Owners                        int    `json:"owners"`
					TotalSalesDecimal             string `json:"totalSalesDecimal"`
					Typename                      string `json:"__typename"`
				} `json:"metrics"`
				Typename string `json:"__typename"`
			} `json:"collection"`
			Typename string `json:"__typename"`
		} `json:"public"`
	} `json:"data"`
}
type NftData struct {
	Data struct {
		Public struct {
			Assets []struct {
				ID                  string `json:"id"`
				Name                string `json:"name"`
				Copies              int    `json:"copies"`
				CopiesInCirculation int    `json:"copiesInCirculation"`
				Creator             struct {
					UUID        string `json:"uuid"`
					ID          string `json:"id"`
					Username    string `json:"username"`
					DisplayName string `json:"displayName"`
					IsCreator   bool   `json:"isCreator"`
					Avatar      struct {
						URL      string `json:"url"`
						Typename string `json:"__typename"`
					} `json:"avatar"`
					IsCreationWithdrawalBlocked      bool        `json:"isCreationWithdrawalBlocked"`
					CreationWithdrawalBlockExpiredAt interface{} `json:"creationWithdrawalBlockExpiredAt"`
					Verified                         bool        `json:"verified"`
					Typename                         string      `json:"__typename"`
				} `json:"creator"`
				Main struct {
					URL      string `json:"url"`
					Typename string `json:"__typename"`
				} `json:"main"`
				Cover struct {
					URL      string `json:"url"`
					Typename string `json:"__typename"`
				} `json:"cover"`
				RoyaltiesRateDecimal   string `json:"royaltiesRateDecimal"`
				PrimaryListingsCount   int    `json:"primaryListingsCount"`
				SecondaryListingsCount int    `json:"secondaryListingsCount"`
				PrimarySalesCount      int    `json:"primarySalesCount"`
				TotalSalesDecimal      string `json:"totalSalesDecimal"`
				DefaultListing         struct {
					EditionID      string `json:"editionId"`
					PriceDecimal   string `json:"priceDecimal"`
					Mode           string `json:"mode"`
					AuctionHasBids bool   `json:"auctionHasBids"`
					Typename       string `json:"__typename"`
				} `json:"defaultListing"`
				DefaultAuctionListing interface{} `json:"defaultAuctionListing"`
				DefaultSaleListing    struct {
					EditionID    string `json:"editionId"`
					PriceDecimal string `json:"priceDecimal"`
					Mode         string `json:"mode"`
					Typename     string `json:"__typename"`
				} `json:"defaultSaleListing"`
				DefaultPrimaryListing   interface{} `json:"defaultPrimaryListing"`
				DefaultSecondaryListing struct {
					EditionID      string `json:"editionId"`
					PriceDecimal   string `json:"priceDecimal"`
					Mode           string `json:"mode"`
					AuctionHasBids bool   `json:"auctionHasBids"`
					Typename       string `json:"__typename"`
				} `json:"defaultSecondaryListing"`
				DefaultSecondaryAuctionListing interface{} `json:"defaultSecondaryAuctionListing"`
				DefaultSecondarySaleListing    struct {
					EditionID    string `json:"editionId"`
					PriceDecimal string `json:"priceDecimal"`
					Mode         string `json:"mode"`
					Typename     string `json:"__typename"`
				} `json:"defaultSecondarySaleListing"`
				Likes               int         `json:"likes"`
				Views               int         `json:"views"`
				IsCurated           bool        `json:"isCurated"`
				DefaultEditionID    string      `json:"defaultEditionId"`
				IsLiked             bool        `json:"isLiked"`
				IsAcceptingOffers   bool        `json:"isAcceptingOffers"`
				IsExternalNft       bool        `json:"isExternalNft"`
				ExternalNftMetadata interface{} `json:"externalNftMetadata"`
				Typename            string      `json:"__typename"`
			} `json:"assets"`
			Typename string `json:"__typename"`
		} `json:"public"`
	} `json:"data"`
}
type Data struct {
	Results []struct {
		Id               string  `json:"id"`
		Price            float64 `json:"price"`
		EscrowPubkey     string  `json:"escrowPubkey"`
		Owner            string  `json:"owner"`
		CollectionName   string  `json:"collectionName"`
		CollectionTitle  string  `json:"collectionTitle"`
		Img              string  `json:"img"`
		Title            string  `json:"title"`
		Content          string  `json:"content"`
		ExternalURL      string  `json:"externalURL"`
		PropertyCategory string  `json:"propertyCategory"`
		Creators         []struct {
			Address  string `json:"address"`
			Verified int    `json:"verified"`
			Share    int    `json:"share"`
		} `json:"creators"`
		SellerFeeBasisPoints int    `json:"sellerFeeBasisPoints"`
		MintAddress          string `json:"mintAddress"`
		Attributes           []struct {
			TraitType string `json:"trait_type"`
			Value     string `json:"value"`
		} `json:"attributes"`
	} `json:"results"`
}
