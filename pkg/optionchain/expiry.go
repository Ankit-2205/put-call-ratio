package optionchain

import (
	"math"
	"sort"

	"github.com/Ankit-2205/put-call-ratio/pkg/repository"
)

func getRecordsPerExpiry(optionChain repository.OptionChain, numOfRecords int) map[string][]repository.RecordData {

	recordsPerExpiry := make(map[string][]repository.RecordData)

	for _, record := range optionChain.Records.Data {
		records, ok := recordsPerExpiry[record.ExpiryDate]
		if !ok {
			records = make([]repository.RecordData, 0)
		}

		records = append(records, record)
		recordsPerExpiry[record.ExpiryDate] = records
	}

	for expiryDate, records := range recordsPerExpiry {

		filteredRecords := []repository.RecordData{}
		sort.Slice(records, func(i, j int) bool {
			return records[i].StrikePrice < records[i].StrikePrice
		})

		highPoint := sort.Search(len(records), func(i int) bool {
			return records[i].StrikePrice > optionChain.Records.UnderlyingValue
		})

		index := highPoint - 1
		toIndex := int(math.Max(0, float64(highPoint)-float64(numOfRecords)))
		for index >= toIndex {
			filteredRecords = append(filteredRecords, records[index])
			index--
		}

		index = highPoint
		toIndex = int(math.Min(float64(len(records)), float64(highPoint)+float64(numOfRecords-1)))
		for index <= toIndex {
			filteredRecords = append(filteredRecords, records[index])
			index++
		}

		recordsPerExpiry[expiryDate] = filteredRecords
	}

	return recordsPerExpiry
}
