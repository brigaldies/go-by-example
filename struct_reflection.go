package main

import (
	"fmt"
)

type GapAnalyzerConfig struct {
	ESPrismGapAnalyzerStoreIDs                        string `config:"alias=es_prism_gap_analyzer_store_ids"` // comma separated store ids
	ESPrismGapAnalyzerBatchSize                       int    `config:"alias=es_prism_gap_analyzer_batch_size, default=400"`
	ESPrismGapAnalyzerDBURL                           string `config:"alias=es_prism_gap_analyzer_db_url"`
	ESPrismGapAnalyzerCatalogIndexURL                 string `config:"alias=es_prism_gap_analyzer_catalog_index_url"`
	ESPrismGapAnalyzerEnableIndexMissing              bool   `config:"alias=es_prism_gap_analyzer_enable_reindex"`
	ESPrismGapAnalyzerPrismProductsUrl                string `config:"alias=es_prism_gap_analyzer_prism_products_url"`
	ESPrismGapAnalyzerReportingUrl                    string `config:"alias=es_prism_gap_analyzer_reporting_url"`
	ESPrismGapAnalyzerEnableReporting                 bool   `config:"alias=es_prism_gap_analyzer_enable_reporting,default=true"`
	ESPrismGapAnalyzerBatchIndexMaxProds              int    `config:"alias=es_prism_gap_analyzer_batch_index_batch_size,default=100"`
	ESPrismGapAnalyzerBatchIndexMaxWaitSeconds        int    `config:"alias=es_prism_gap_analyzer_batch_index_max_wait_seconds,default=1800"`
	ESPrismGapAnalyzerIndexingMode                    string `config:"alias=es_prism_gap_analyzer_indexing_mode,default=batch"`
	ESPrismGapAnalyzerIndexingTimeoutSeconds          int    `config:"alias=es_prism_gap_analyzer_indexing_timeout_seconds,default=45"` // large enough for batch i hope
	ESPrismGapAnalyzerLogItemsSummaryAt               int    `config:"alias=es_prism_gap_analyzer_log_item_summary_at,default=5000"`
	ESPrismGapAnalyzerVerifyWithPrism                 bool   `config:"alias=es_prism_gap_analyzer_verify_with_prism,default=true"`
	ESPrismGapAnalyzerPrometheusURL                   string `config:"alias=es_prism_gap_analyzer_prometheus_url"`
	ESPrismGapAnalyzerPrometheusCheckMinutes          int    `config:"alias=es_prism_gap_analyzer_prometheus_check_minutes,default=10"`
	ESPrismGapAnalyzerMaxLagBeforeProcess             int    `config:"alias=es_prism_gap_analyzer_max_lag_before_process,default=500000"`
	ESPrismGapAnalyzerMaxWaitMinutesForLagBeforeExit  int    `config:"alias=es_prism_gap_analyzer_max_wait_minutes_for_lag_before_exit,default=360"`
	EsPrismGapAnalyzerReportingMaxRetries             int    `config:"alias=es_prism_gap_analyzer_reporting_max_retries,default=3"`
	EsPrismGapAnalyzerDBQueryTimeRange                string `config:"alias=es_prism_gap_analyzer_db_query_time_range"`                                  // 2 timestamps separated by 'to' i.e: 2022-12-19 13:00:00 to 2022-12-21 18:00:00
	EsPrismGapAnalyzerDBQueryTimeRangeUpToPreviousDay bool   `config:"alias=es_prism_gap_analyzer_db_query_time_range_up_to_previous_day,default=false"` // Used for automatic and daily exec. of the job
	ESPrismGapAnalyzerDBQueryOrderBy                  string `config:"alias=es_prism_gap_analyzer_db_query_order_by,default="`
}

type LTRConfiguration struct {
	LTRExperiments            string `config:"ltr_experiments"`
	RescoreModels             string `config:"rescore_models"`
	RescoreWindowSize         int    `config:"rescorer_window_size,default=1000"`
	RescoreQueryWeight        int    `config:"rescore_query_weight,default=0"`
	AdsRescoreQueryWeight     int    `config:"ads_rescore_query_weight,default=10"`
	RescoreModelWeight        int    `config:"rescore_model_weight,default=1"`
	LTRDataCollectionEnabled  bool   `config:"alias=ltr_data_collection_enabled"`
	LTRVariationBucketSize    int    `config:"alias=ltr_variation_bucket_size, default=300"`
	LTRTotalVariationSize     int    `config:"alias=ltr_total_variation_size, default=2000"`
	ProductsReScoreSize       int    `config:"products_re_score_size,default=48"`
	VectorSearchEnabledStores string `config:"vector_search_enabled_stores,default=27"` // array of values separated by ','
	Phase1RankerExperiments   string `config:"phase1_ranker_experiments"`
	Phase1RankerScript        string `config:"phase1_ranker_script,default=shipt-search-ranker-v2.2"`
	LTREmbeddingVersion       string `config:"ltr_embedding_version,default=V2"`
}

type Configuration struct {
	GapAnalyzerConfig
	LTRConfiguration
	Foo string
	Bar int
}

func main() {
	cfg := Configuration{}
	fmt.Println(fmt.Sprintf("%+v", cfg))

	//t := reflect.TypeOf(cfg)
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Println(t.Field(i).Name)
	//}
}
