package main

import (
	"context"
	"encoding/json"
	"fmt"
	prometheus "github.com/prometheus/client_golang/api"
	prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"os"
	"strings"
	"time"
)

const (
	//  query = "sum(kafka_consumergroup_lag{topic=~\"prd.pipeline.dsom.catalog\", consumergroup=~\"prd.catalog-worker\"}) by (consumergroup, topic)"
	query = "sum(kafka_consumergroup_lag{topic=\"prd.pipeline.dsom.catalog\", consumergroup=\"prd.catalog-worker\"}) by (consumergroup)"
	//query         = "kafka_consumergroup_lag{topic=\"prd.pipeline.dsom.catalog\", consumergroup=\"prd.catalog-worker\"}"
	//prometheusUrl = "https://prometheus.us-central1.shipt.com"
	prometheusUrl = "https://thanos-query.gcp.shipttech.com/"
)

func main() {

	prometheusClient, err := prometheus.NewClient(prometheus.Config{
		Address: prometheusUrl,
	})

	if err != nil {
		fmt.Printf("prometheus.NewClient() failed: %v\n", err)
		os.Exit(1)
	}

	api := prometheus_v1.NewAPI(prometheusClient)

	start := time.Now()
	result, warnings, err := api.Query(context.Background(), query, time.Now(), prometheus_v1.WithTimeout(30*time.Second))
	if warnings != nil {
		fmt.Printf("Query (Now) warnings: %v\n", strings.Join(warnings, ","))
	}
	if err != nil {
		fmt.Printf("Query (Now) error: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("[%v] Query (Now) result: type=%v, value=\"%v\"\n", time.Since(start), result.Type(), result.String())

	marshaled, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Query (Now) result marshal error: %v\n", err)
		os.Exit(3)
	}

	var unmarshaled []*model.Sample // from import "github.com/prometheus/common/model"
	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		fmt.Printf("Query (Now) marshalled result umarshall error: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("Query (Now) result model: %+v\n", unmarshaled)

	timeNow := time.Now()
	timeAnHourAgo := timeNow.Add(-time.Duration(1) * time.Hour)
	fmt.Printf("Collecting metrics between %v and %v\n", timeAnHourAgo, timeNow)
	start = time.Now()
	result, warnings, err = api.QueryRange(
		context.Background(),
		query,
		prometheus_v1.Range{
			Start: timeAnHourAgo,
			End:   timeNow,
			Step:  time.Duration(1) * time.Minute,
		},
		prometheus_v1.WithTimeout(30*time.Second))
	if warnings != nil {
		fmt.Printf("Query (Range) warnings: %v\n", strings.Join(warnings, ","))
	}
	if err != nil {
		fmt.Printf("Query (Range) error: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("[%v] Query (Range) result: type=%v, value=\"%v\"\n", time.Since(start), result.Type(), result.String())

	marshaled, err = json.Marshal(result)
	if err != nil {
		fmt.Printf("Query (Range) result marshal error: %v\n", err)
		os.Exit(3)
	}

	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		fmt.Printf("Query (Range) marshalled result umarshall error: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("Query (Range) result model: %+v\n", unmarshaled)
}
