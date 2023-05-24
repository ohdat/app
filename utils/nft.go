package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TokenInfo struct {
	Image string `json:"image"`
}

func ipfs2http(uri string) string {
	if uri[:6] == "ipfs:/" {
		return "https://ipfs.io/ipfs/" + uri[7:]
	}
	return uri
}

func Uri2Img(uri string) (img string) {
	uri = ipfs2http(uri)
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Printf("http.Get err: %v", err)
		time.Sleep(time.Microsecond * 100)
		return Uri2Img(uri)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 429 {
			time.Sleep(time.Second * 1)
			return Uri2Img(uri)
		}
		fmt.Printf("uri: %s resp.StatusCode: %v \n", uri, resp.StatusCode)
		return ""
		//time.Sleep(time.Microsecond * 100)
		//return Uri2Img(uri)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("uri: %s ioutil.ReadAll err: %v \n", uri, err)
		return ""
	}
	var item TokenInfo
	json.Unmarshal(body, &item)
	img = item.Image
	return ipfs2http(img)
}
