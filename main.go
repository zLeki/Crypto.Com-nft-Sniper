package main

import (
	types "SolanaNftSniper/handler"
	"bytes"
	"encoding/json"
	"fmt"
	cpu2 "github.com/shirou/gopsutil/cpu"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
5/15/2022
In words, we are trying to create a program that scrapes prices from the Solana NFT marketplace.
We are going to be using SolSea.io where nfts are sold there. I have found the network api which is used to scrape the data.
https://api.solanart.io/get_nft?collection=degenape&page=0&limit=20&order=price-DESC&min=0&max=99999&search=&listed=true&fits=all&bid=all is the url to get the data.
We would have to change the collection parameter to the name of the collection we want to scrape found in the json file.

5/18/22
Created a wait group and anonymous go routines to handle the scraping. Additionally, I am going to create a cui to manage price data. And we can now fetch cheap nfts once we want to purchase

5/19/22
Manage price data. Watch cpu usage. Notifiy via discord when to purchase. Calculate percentage. Planning on adding notification to notifiy when to buy when nft is over 50% profit or over any range.
*/

var (
	nftPrices                = make(map[string]int)
	c                        *types.Config
	RecentPullBackPercentage float64
	originalPrices           = make(map[string]int)
	clear                    map[string]func()
	findSerial               = func(name string, c *types.Config) (string, error) {
		for _, v := range c.Projects {
			if strings.ToLower(name) == strings.ToLower(v.Name) {
				return strings.Split(v.Url, "collection/")[1], nil
			}
		}
		return "", fmt.Errorf("could not find serial for %s", name)
	}
)

const (
	Paused    = 1
	Running   = 2
	Stopped   = 3
	confident = "█"
	uncertain = "░"
)

func main() {
	config, _ := ioutil.ReadFile("./config/configuration.json")
	var configData types.Config
	json.Unmarshal(config, &configData)
	c = &configData

	runtime.GOMAXPROCS(runtime.NumCPU())
	go UpdateOriginalPrices()
	var wg sync.WaitGroup
	wg.Add(5 + 1)
	workers := make([]chan int, 5)
	for i := 0; i < 5; i++ {
		workers[i] = make(chan int, 1)
		go func(i int) {
			SnipeThread(c, workers[i])
		}(i)

	}
	go func() {
		Controller(workers)
		wg.Done()
	}()
	go Cui(c)

	wg.Wait()

}
func Controller(workers []chan int) {
	ChangeState(workers, Running)
}
func ChangeState(workers []chan int, state int) {
	for _, worker := range workers {
		worker <- state
	}
}
func SnipeThread(c *types.Config, ws <-chan int) {
	state := Paused

	for {
		select {
		case state = <-ws:
			switch state {
			case Paused:
			//pause
			case Running:
			case Stopped:
				fmt.Println("NFT Sniper Stopped")
				return
			}
		default:
			runtime.Gosched()
			if state == Paused {
				break
			}
			for i, v := range c.GetFloor() {
				nftPrices[i] = int(v)
			}
		}

	}

}
func Cui(c *types.Config) {
	fmt.Println("Starting NFT Sniper")
	fmt.Println("Version: Alpha 1.01")
	fmt.Println("Author: Leki")
	fmt.Println("Email: any-grid05@icloud.com")
	fmt.Println("Website: https://www.leki.sbs/portfolio")
	fmt.Println("Github: github.com/zLeki")
	var mutex = &sync.Mutex{}
	go func() {
		for {
			mutex.Lock()
			for i, v := range nftPrices {
				for o, r := range originalPrices {
					if v < r && o == i {
						log.Println("Before: ", originalPrices[o], "After: ", nftPrices[i])
						log.Println("NFT: ", strings.Title(o), " Price: ", r, " Original Price: ", r, " Floor Price Decreased by: ", r-v, "Now at: ", v, "Loss Percentage: ", (r-v)*100/r)
						serial, err := findSerial(o, c)
						if err != nil {
							log.Fatal(err)
						}
						nfts, _ := c.GetNfts(serial)

						discordWebhookData := []byte(`{"content":"","embeds":[{"title":"Hit","url":"https://crypto.com/nft/collection/` + serial + `?buyNow=true&sort=price&order=ASC&asset=` + nfts.Data.Public.Assets[0].ID + `&edition=` + nfts.Data.Public.Assets[0].DefaultSaleListing.EditionID + `&detail-page=MARKETPLACE","color":5814783,"fields":[{"name":"NFT","value":"` + nfts.Data.Public.Assets[0].Name + `","inline":true},{"name":"New Price","value":"$` + strconv.Itoa(v) + `","inline":true},{"name":"Original Price","value":"$` + strconv.Itoa(r) + `","inline":true},{"name":"Profit", "value": "` + strconv.FormatFloat(float64(r-v)*100/float64(r), 'f', 2, 64) + `%"}],"image":{"url":"` + nfts.Data.Public.Assets[0].Cover.URL + `"}}],"attachments":[]}`)
						req, _ := http.Post("https://canary.discord.com/api/webhooks/976650669563981864/f9UX6JHyc5odpq4DojpIJxYLVPEFuBqyJQGf0pAWkg4eyXZM1fPBpDU2hrWhGDC-HLiQ", "application/json", bytes.NewBuffer(discordWebhookData))
						if req.StatusCode != 200 {
							log.Println("Discord Webhook Error: ", req.StatusCode)
						}

						UpdateOriginalPrices()
					}
				}
			}

			mutex.Unlock()
			time.Sleep(time.Millisecond * 10)

		}

	}()
	for {

		fmt.Println(
			"\n                       __                  _   _             _               _            _   _      _      \n                      / _|                | | (_)           (_)             | |          | | (_)    | |     \n  _ __   ___  _ __   | |_ _   _ _ __   ___| |_ _  ___  _ __  _ _ __   __ _  | |_ ___  ___| |_ _  ___| | ___ \n | '_ \\ / _ \\| '_ \\  |  _| | | | '_ \\ / __| __| |/ _ \\| '_ \\| | '_ \\ / _` | | __/ _ \\/ __| __| |/ __| |/ _ \\\n | | | | (_) | | | | | | | |_| | | | | (__| |_| | (_) | | | | | | | | (_| | | ||  __/\\__ \\ |_| | (__| |  __/\n |_| |_|\\___/|_| |_| |_|  \\__,_|_| |_|\\___|\\__|_|\\___/|_| |_|_|_| |_|\\__, |  \\__\\___||___/\\__|_|\\___|_|\\___|\n                                                                      __/ |                                 \n                                                                     |___/                                  \n")
		mutex.Lock()
		for i, v := range nftPrices {
			log.Println("NFT: ", strings.Title(i), " Price: ", v)
		}
		mutex.Unlock()
		cpuUsage, err := cpu2.Percent(0, false)
		if err != nil {
			log.Fatal(err)
		}
		confidence := int(cpuUsage[0])
		fmt.Printf("CPU Usage [")
		for i := 0; i < confidence/10; i++ {
			fmt.Printf(confident)
		}
		for i := 0; i < 10-confidence/10; i++ {
			fmt.Printf(uncertain)
		}
		fmt.Printf("]\n")
		time.Sleep(time.Second * 1)

		clear[runtime.GOOS]()
	}

}
func UpdateOriginalPrices() bool {
	for i, v := range c.GetFloor() {
		originalPrices[i] = int(v)
	}
	return true
}

func init() {
	clear = make(map[string]func())
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
