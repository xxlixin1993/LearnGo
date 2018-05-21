// 抓取tripadvisor网站游记
package main

import (
	"net/http"
	"fmt"
	"strconv"
	"time"
	"log"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic"
	"context"
	"os"
)

var (
	// 要抓取的游记最大id
	tripadvisorTotalId = 3

	// 起多少个goroutine去抓取
	fetchGoroutineTotal = 3

	// 要抓取的url ex: https://www.tripadvisor.cn/TourismBlog-t6598
	tripadvisorDetail = "https://www.tripadvisor.cn/TourismBlog-t"

	// es client
	esClient *elastic.Client

	tPool map[int]*Tripadvisor
)

// es 索引
const ktripadvisorTitleIndex = "tti"

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

	var esErr error
	esClient, esErr = elastic.NewClient()
	if esErr != nil {
		log.Fatalf("es client err : %s", esErr)
		os.Exit(10)
	}

	doTripadvisor()

	secs := time.Since(start).Seconds()

	fmt.Printf("time: %f", secs)
}

func newTripadvisor() *Tripadvisor {
	return &Tripadvisor{
		urlChan: make(chan string),
		done:    make(chan int),
	}
}

func newEsChannel() *EsChannel {
	return &EsChannel{
		esChan: make(chan *EsContent),
		done:   make(chan int),
	}
}

// 开始获取页面信息
func doTripadvisor() {
	tPool = make(map[int]*Tripadvisor)

	esChan := newEsChannel()

	esChan.esg.Add(1)
	go esChan.output()

	for gnum := 0; gnum < fetchGoroutineTotal; gnum ++ {
		tPool[gnum] = newTripadvisor()

		tPool[gnum].twg.Add(1)
		go tPool[gnum].fetchTripadvisor(esChan)
	}

	for i := 1; i <= tripadvisorTotalId; i++ {
		tPool[i%fetchGoroutineTotal].urlChan <- tripadvisorDetail + strconv.Itoa(i)
	}

	for key := range tPool {
		close(tPool[key].done)
		tPool[key].twg.Wait()
	}

	close(esChan.done)
	esChan.esg.Wait()
}

// 写入es
func (esc *EsChannel) output() {
	defer esc.esg.Done()
	for {
		select {
		case <-esc.done:
			close(esc.esChan)
			return
		case data := <-esc.esChan:
			// 判断必须有title才能输出到es
			// 需要先建es index和中文分词option
			// 1. curl -XPUT http://localhost:9200/tti
			// 2. curl -XPOST http://localhost:9200/tti/fulltext/_mapping -H 'Content-Type:application/json' -d'
			//{
			// "properties": {
			//     "content": {
			//         "type": "text",
			//         "analyzer": "ik_max_word",
			//         "search_analyzer": "ik_max_word"
			//     }
			// }
			//}'
			if data.Title != "" {
				put1, err := esClient.Index().Index(ktripadvisorTitleIndex).Type("fulltext").
					BodyJson(data).Do(context.Background())

				if err != nil {
					panic(err)
				}
				fmt.Printf("Indexed tti  %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
			}
		}
	}
}

// 抓取
func (t *Tripadvisor) fetchTripadvisor(esChan *EsChannel) {
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
			esChan.esChan <- esContent
		}
	}
}