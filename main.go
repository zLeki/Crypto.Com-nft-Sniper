package main

import (
	types "SolanaNftSniper/handler"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
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
*/

var (
	nftPrices                = make(map[string]int)
	c                        *types.Config
	RecentPullBackPercentage float64
	originalPrices           = make(map[string]int)
	clear                    map[string]func()
)

const (
	Paused  = 1
	Running = 2
	Stopped = 3
)

func main() {
	config, _ := ioutil.ReadFile("./config/configuration.json")
	var configData types.Config
	json.Unmarshal(config, &configData)
	c = &configData
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i, v := range c.GetFloor() {
		originalPrices[i] = int(v)
		fmt.Println(i, v)
		fmt.Println(originalPrices)
	}
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
	for {

		for i, v := range nftPrices {
			for o, r := range originalPrices {
				if v < r && o == i {
					log.Println("Before: ", originalPrices[o], "After: ", nftPrices[i])
					log.Println("NFT: ", strings.Title(o), " Price: ", r, " Original Price: ", v, " Floor Price Decreased by: ", r-v, "Now at: ", v, "Loss Percentage: ", float64(r-v)/float64(r)*100)
				}
			}
			log.Println("NFT: ", strings.Title(i), " Price: ", v, "Overall Decrease since start:", originalPrices[i]-v)
		}
		time.Sleep(1 * time.Second)
		clear[runtime.GOOS]()
	}
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
