package models

type RequestStats struct {
	totalRequests      uint64
	successfulRequests uint64
	failedRequests     uint64
	totalLatency       int64
	maxConcurrent      int64
	currentConcurrent  int64
}
