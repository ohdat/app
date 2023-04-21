package nft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type TransferRes struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Transfers []struct {
			Blocknum        string      `json:"blockNum"`
			Uniqueid        string      `json:"uniqueId"`
			Hash            string      `json:"hash"`
			From            string      `json:"from"`
			To              string      `json:"to"`
			Value           interface{} `json:"value"`
			Erc721Tokenid   string      `json:"erc721TokenId"`
			Erc1155Metadata interface{} `json:"erc1155Metadata"`
			Tokenid         string      `json:"tokenId"`
			Asset           interface{} `json:"asset"`
			Category        string      `json:"category"`
			Rawcontract     struct {
				Value   interface{} `json:"value"`
				Address string      `json:"address"`
				Decimal interface{} `json:"decimal"`
			} `json:"rawContract"`
		} `json:"transfers"`
		Pagekey string `json:"pageKey"`
	} `json:"result"`
}

type GetNFT struct {
	Ownednfts []struct {
		Contract struct {
			Address string `json:"address"`
		} `json:"contract"`
		ID struct {
			Tokenid string `json:"tokenId"`
		} `json:"id"`
		Balance string `json:"balance"`
	} `json:"ownedNfts"`
	Pagekey    string `json:"pageKey"`
	Totalcount int    `json:"totalCount"`
	Blockhash  string `json:"blockHash"`
}

func GetNft() *Nft {
	var key = viper.GetString("alchemy.key")
	if key == "" {
		panic("alchemy.key is empty")
	}
	return NewNft(key)
}

func NewNft(key string) *Nft {
	return &Nft{
		Key: key,
	}
}

type Nft struct {
	Key string
}

func (s Nft) Total(owner string) (total int, err error) {

	url := fmt.Sprintf("https://eth-mainnet.g.alchemy.com/nft/v2/%s/getNFTs?owner=%s&pageSize=10&withMetadata=false", s.Key, owner)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var item GetNFT
	err = json.Unmarshal(body, &item)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(body))
	return item.Totalcount, nil
}
func (s Nft) TransferTotal(owner string) (total int) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		count, _ := s.ToTransfer(owner, "")
		total += count
		log.Println("ToTransfer:", count)
		wg.Done()
	}()
	go func() {
		count, _ := s.FormTransfer(owner, "")
		total += count
		log.Println("FormTransfer:", count)
		wg.Done()
	}()
	wg.Wait()
	return total
}
func (s Nft) FormTransfer(owner string, pageKey string) (total int, err error) {

	url := fmt.Sprintf("https://eth-mainnet.g.alchemy.com/v2/%s", s.Key)
	var payload *strings.Reader
	if pageKey == "" {
		payload = strings.NewReader(fmt.Sprintf(`{
		 "id": 1,
		 "jsonrpc": "2.0",
		 "method": "alchemy_getAssetTransfers",
		 "params": [
			  {
				   "fromBlock": "0x0",
				   "toBlock": "latest",
				   "category": [
						"erc721",
						"erc1155"
				   ],
				   "withMetadata": false,
				   "excludeZeroValue": true,
				   "maxCount": "0x3e8",
				   "fromAddress": "%s"
			  }
		 ]
	}`, owner))
	} else {
		payload = strings.NewReader(fmt.Sprintf(`{
		 "id": 1,
		 "jsonrpc": "2.0",
		 "method": "alchemy_getAssetTransfers",
		 "params": [
			  {
				   "fromBlock": "0x0",
				   "toBlock": "latest",
				   "category": [
						"erc721",
						"erc1155"
				   ],
				   "withMetadata": false,
				   "excludeZeroValue": true,
				   "maxCount": "0x3e8",
				   "fromAddress": "%s",
				   "pageKey": "%s"
			  }
		 ]
	}`, owner, pageKey))
	}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Transfer,Err", err)
		return 0, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Transfer,Err", err)
		return 0, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(string(body))
	var resp TransferRes
	json.Unmarshal(body, &resp)
	fmt.Println("LEN", len(resp.Result.Transfers))
	if resp.Result.Pagekey != "" {
		total, err = s.FormTransfer(owner, resp.Result.Pagekey)
		return total + len(resp.Result.Transfers), err
	}
	return len(resp.Result.Transfers), nil
}
func (s Nft) ToTransfer(owner string, pageKey string) (total int, err error) {

	url := fmt.Sprintf("https://eth-mainnet.g.alchemy.com/v2/%s", s.Key)
	var payload *strings.Reader
	if pageKey == "" {
		payload = strings.NewReader(fmt.Sprintf(`{
			"id": 1,
			"jsonrpc": "2.0",
			"method": "alchemy_getAssetTransfers",
			"params": [
				{
					"fromBlock": "0x0",
					"toBlock": "latest",
					"category": [
							"erc721",
							"erc1155"
					],
					"withMetadata": false,
					"excludeZeroValue": true,
					"maxCount": "0x3e8",
					"toAddress": "%s"
				}
			]
		}`, owner))
	} else {
		payload = strings.NewReader(fmt.Sprintf(`{
			"id": 1,
			"jsonrpc": "2.0",
			"method": "alchemy_getAssetTransfers",
			"params": [
				{
					"fromBlock": "0x0",
					"toBlock": "latest",
					"category": [
							"erc721",
							"erc1155"
					],
					"withMetadata": false,
					"excludeZeroValue": true,
					"maxCount": "0x3e8",
					"toAddress": "%s",
					"pageKey": "%s"
				}
			]
		}`, owner, pageKey))
	}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Transfer,Err", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Transfer,Err", err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(string(body))
	var resp TransferRes
	json.Unmarshal(body, &resp)
	fmt.Println("LEN", len(resp.Result.Transfers))
	if resp.Result.Pagekey != "" {
		total, err = s.ToTransfer(owner, resp.Result.Pagekey)
		return total + len(resp.Result.Transfers), err
	}
	return len(resp.Result.Transfers), nil
}
