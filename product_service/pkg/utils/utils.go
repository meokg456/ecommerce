package utils

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func SplitDynamoBatchRequest[T any](items []T, tableName string) ([]dynamodb.BatchWriteItemInput, error) {
	maxItems := 25
	length := len(items)
	totalRequests := length / maxItems
	remainder := length % maxItems

	var batchWriteItemInputs []dynamodb.BatchWriteItemInput

	for i := range totalRequests + 1 {
		n := maxItems
		if i == totalRequests {
			n = remainder
		}
		start := i * maxItems

		var writeReqs []types.WriteRequest

		for j := range n {
			item := items[start+j]
			av, err := attributevalue.MarshalMap(item)
			if err != nil {
				return nil, err
			}
			writeReqs = append(writeReqs, types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: av,
				},
			})

		}

		batchWriteItemInputs = append(batchWriteItemInputs, dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				tableName: writeReqs,
			},
		})
	}

	return batchWriteItemInputs, nil
}
