package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)


//一分钟阅读红包 总1800
func main()  {
	requestKey := [] string {
		"params=H0QQESkaBgZQX1Y2AxAZNlNQQyIJARdEWFYrADoQFxcCARRWTn1BVVJLVUxWU1RGb0lRUkhRRUlVVlh9AAgnFxMdBwM9EH1JR1FCVE1UXkVGbUNQUkJURlRVQUw6RFBTSgBEUFRGRGpDAVZLU0ZWUk1Gb0BdUkZVRQVVQRZqEgBUQVVNVldFVnNRCgYfR05GIjIgC1FJQRsWAw0AHVZlURIKFAxWSEQdGTpRX0FKU0xWUkVEbEJQW0ZdRlFEWFYzB0dZUAYeIhwnIQUiKyZDVzkjDgc5bCkqMAgsRAc%2BNicFClwLEVYuCT8YPg48MFZDMjEtVy0eMzAyDzwnESMwMiIYKQ0yKisNNgtFOT4KXFUTICImPEQcFT4yDzwEHi4JIiMzHQYJHhwiDihFLjUfNDpCXBoqIRdNDyJYXlBJVhcJEgArChUGUF9WMAkBIDYSCkFeRxcNEg1WZVGA2s2Aw%2FpEWFYvHBYKBgwbCkROVrrK2ofK%2BVZIRBsHfUlHIhwBBgsPEFZzURQKFkdORh4dFTAeDFJLVUxWU1ZYfQcWQUhHRVFQQkNsRFRbRUdYRgcaEC0cDAc7AVZeRBJBO0cGVERQR1MHREI%2BRlJBXkcHCwAAGj4eAEFIRzAiMiA1MRcXDBsBVkhEFwU2F0dZUB0dBQkZHX1fRwAGFlZeV0FCaURWVENdQ1ZUTFh9BwwOFxcrEB8EEX1JRw0XEgc7Eh0ZOgFHT1AEFwcPEFZlUV1XSlxNU1VGTX1fRwwBOgIBFAcdMB1HWVAkGgAUGx07REtSXFdWSEQQESkaBgYtBwYFCBBWZVEdChMKGQ1EWFYxHAEGAUdORj0PKH0HHBMXOVZeOlYaOgQWP1BJKEYIGxA6L0dZREkoRhMGGDkBCg4uR044RAAbKgcMAh05Vkg6VgEtHxEMLkdOOEQcACsDFlkuOShLOigocB5UTQYRWgcIKCgDXAgMEAwYATooKHBCXFNKV0BUX0ZHbktWUkRLHBALGCh9DkkYLkcAHRYRKH1JOUEcAAMXOlZYA1ELDBYAKEZcRUJzL0cWAAkSFgkZKH1JOUEGCgEQDxUbA1FJP1AQBggSGyh9STlBGhEAFBVOKAMvSj8uOVsJV1oAK10GDS45KEsLGxY2HwA%2FLjlbVV9ETG1HVVpAVkVcVUVCcRsRDh45VhlKDyh9BxwTFzlWXjpWGjoEFj9QSShGCBsQOi9HWUBTWDhEAQYzFRcMHzlWXjpWADAGEQoTCihGSihWKgEJFx05Vl46VhwrBxUQSDkoOEkoKANcCFJcEQBKBRooAy9KDh0HHQgDKCgDXFRaQl1GUFZNRmxCXVBDU1oMEhkYA1EYTwk5VhAfBBEDUV8%2FUAsRExU1ECkvR08uRxoLAhEofUlWU145VhEUGBItHAg%2FUF8oRhIbASsaBAwuR1g4RAEGMwcKP1BfKEYEFR07BhYHGTlWGTtWWH0FABFQX1ZWSEFaalEY",
	}

	for _,v :=range requestKey {
		for i := 0; i < 50; i++ {
			httpDo("",v)
			time.Sleep(time.Duration(20)*time.Second)
		}
	}

}

func httpDo(urls string , s string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://yuedu4.dftoutiao.com/index/Yuedutimer/read_news", strings.NewReader(s))
	if err != nil {
		// handle error
	}


	req.Header.Set("POST", "/index/Yuedutimer/read_news HTTP/1.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", "1497")
	req.Header.Set("Host", "yuedu4.dftoutiao.com")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "okhttp/3.4.1")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("ok")
}