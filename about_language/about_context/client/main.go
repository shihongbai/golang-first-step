package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type resData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives: true,
	}

	client := http.Client{
		Transport: &transport,
	}

	respChan := make(chan *resData, 1)
	req, err := http.NewRequest("GET", "https://127.0.0.8000/", nil)
	if err != nil {
		fmt.Println("new request error", err)
		return
	}

	req = req.WithContext(ctx)
	var wg sync.WaitGroup

	wg.Add(1)
	defer wg.Wait()

	go func() {
		resp, err := client.Do(req)

		fmt.Println("resp err", err)
		r := &resData{
			resp: resp,
			err:  err,
		}

		respChan <- r
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		// call api timeout
		fmt.Println("ctx done")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Println("call server api err", result.err)
			return
		}

		defer result.resp.Body.Close()
		data, _ := io.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // 调用cancel释放子goroutine资源
	doCall(ctx)
}
