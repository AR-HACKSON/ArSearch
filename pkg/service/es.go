package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"

	"ArSearch/pkg/service/service_schema"
)

var es *elasticsearch.Client
var bi esutil.BulkIndexer

func init() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://45.76.151.181:9200",
			"http://45.76.151.181:9201",
		},
	}

	es, _ = elasticsearch.NewClient(cfg)

	bi, _ = esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "mirror_search",  // The default index name
		Client:        es,               // The Elasticsearch client
		NumWorkers:    10,               // The number of worker goroutines
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
}

func PutToEs(article *service_schema.ArArticle) (string, error) {
	m1, _ := json.Marshal(article)

	req := esapi.IndexRequest{
		Index:        "ar_search",
		DocumentType: "ar_article",
		DocumentID:   article.ID,
		Body:         strings.NewReader(string(m1)),
		Refresh:      "true",
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func SaveMirrorData(mirrorData *service_schema.MirrorData) (string, error) {
	m1, _ := json.Marshal(mirrorData)

	req := esapi.IndexRequest{
		Index:        "mirror_search_v1",
		DocumentType: "mirror_article",
		DocumentID:   mirrorData.ArweaveTx,
		Body:         strings.NewReader(string(m1)),
		Refresh:      "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func SearchInEs(termQuery string) ([]service_schema.ArSearchRes, error) {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{"match": map[string]interface{}{"article_context": termQuery}},
					{"match": map[string]interface{}{"title": termQuery}},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("ar_search"),
		es.Search.WithDocumentType("ar_article"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		return []service_schema.ArSearchRes{}, err
	}
	defer res.Body.Close()

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	searchResList := make([]service_schema.ArSearchRes, 0)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		score := hit.(map[string]interface{})["_score"]
		res := service_schema.ArSearchRes{
			Score: score.(float64),
			Article: service_schema.ArArticle{
				ID:             article["id"].(string),
				ArticleContext: article["article_context"].(string),
				Title:          article["title"].(string),
			},
			RedirectUrl: fmt.Sprintf("https://arweave.net/%s", article["id"].(string)),
		}

		searchResList = append(searchResList, res)
	}

	return searchResList, nil
}

func SearchMirrorData(termQuery string) ([]service_schema.MirrorSearchRes, error) {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{"match": map[string]interface{}{"content": termQuery}},
					{"match": map[string]interface{}{"title": termQuery}},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("mirror_search_v1"),
		es.Search.WithDocumentType("mirror_article"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithSize(30),
	)

	if err != nil {
		return []service_schema.MirrorSearchRes{}, err
	}
	defer res.Body.Close()

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	searchResList := make([]service_schema.MirrorSearchRes, 0)

	distinctMap := make(map[string]bool, 30)

	if r["hits"] == nil {
		log.Println("hits nil")
		return []service_schema.MirrorSearchRes{}, nil
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		keys := make([]string, 0)
		for k, _ := range article {
			keys = append(keys, k)
		}

		uniqueId := article["originalDigest"].(string)

		if _, ok := distinctMap[uniqueId]; ok {
			continue
		}
		distinctMap[uniqueId] = true

		score := hit.(map[string]interface{})["_score"]

		//todo 这里处理有点问题
		//"2022-05-11T23:58:06+08:00"
		createAt := article["createdAt"].(string)
		createAt = createAt[:10]
		t1, _ := time.Parse("2006-01-02", createAt)

		searchRes := service_schema.MirrorSearchRes{
			MirrorData: service_schema.MirrorData{
				//Id:              article["id"].(int64),
				Title:           article["title"].(string),
				Content:         article["content"].(string),
				CreatedAt:       t1,
				PublishedAt:     t1,
				Digest:          article["digest"].(string),
				Link:            article["link"].(string),
				OriginalDigest:  article["originalDigest"].(string),
				PublicationName: article["publicationName"].(string),
				Cursor:          article["cursor"].(string),
				ArweaveTx:       article["arweaveTx"].(string),
				//BlockHeight:     article["blockHeight"].(float64),
			},
			ArweaveLink: fmt.Sprintf("https://arweave.net/%s", article["arweaveTx"].(string)),
			Score:       score.(float64),
		}

		searchResList = append(searchResList, searchRes)
	}

	return searchResList, nil
}
