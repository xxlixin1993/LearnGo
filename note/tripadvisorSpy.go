// tripadvisor
package main

import (
	"net/http"
	"fmt"
	"strconv"
	"time"
	"log"
	"sync"
	"github.com/PuerkitoBio/goquery"
)

// 要抓取的url ex: https://www.tripadvisor.cn/TourismBlog-t6598
var tripadvisorDetail = "https://www.tripadvisor.cn/TourismBlog-t"
// 要抓取的游记最大id
var tripadvisorTotal = 3
// 起多少个goroutine
var goroutineTotal = 1

type Tripadvisor struct {
	urlChan chan string
	done    chan int
	twg     sync.WaitGroup
}

type EsContent struct {
	Title   string
	Url     string
	Content string
}

type EsChannel struct {
	esChan chan *EsContent
	done   chan int
	esg    sync.WaitGroup
}

func main() {
	start := time.Now()

	doTripadvisor()

	secs := time.Since(start).Seconds()

	fmt.Printf("time: %f", secs)
}

func doTripadvisor() {
	t := &Tripadvisor{
		urlChan: make(chan string),
		done:    make(chan int),
	}

	esc := &EsChannel{
		esChan: make(chan *EsContent),
		done:   make(chan int),
	}
	esc.esg.Add(1)
	go esc.output()

	for gnum := 0; gnum < goroutineTotal; gnum ++ {
		t.twg.Add(1)
		go t.fetchTripadvisor(esc)
	}

	for i := 1; i <= tripadvisorTotal; i++ {
		t.urlChan <- tripadvisorDetail + strconv.Itoa(i)
	}

	close(t.done)
	t.twg.Wait()
	close(esc.done)
	esc.esg.Wait()
}

func (esc *EsChannel) output() {
	defer esc.esg.Done()
	for {
		select {
		case <-esc.done:
			close(esc.esChan)
			return
		case <-esc.esChan:
			// TODO output es
		}
	}

}

func (t *Tripadvisor) fetchTripadvisor(esc *EsChannel) {
	defer t.twg.Done()
	for {
		select {
		case <-t.done:
			close(t.urlChan)
			return
		case url := <-t.urlChan:
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
				return
			}
			//body, err := ioutil.ReadAll(resp.Body)
			//fmt.Println(string(body))
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			resp.Body.Close()

			if err != nil {
				log.Fatal(err)
			}

			title := doc.Find(".title-text").Text()

			s := doc.Find(".strategy-description").Each(func(i int, s *goquery.Selection) {

			})
			esContent := &EsContent{
				Title:   title,
				Content: s.Text(),
				Url:     url,
			}
			fmt.Println(esContent)
			esc.esChan <- esContent
		}
	}
}
