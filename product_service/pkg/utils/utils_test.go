package utils_test

import (
	"testing"

	"github.com/meokg456/productservice/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSplitDynamoBatchRequest(t *testing.T) {
	items := []int{}

	for i := range 95 {
		items = append(items, i)
	}

	batchWriteItemInputs, err := utils.SplitDynamoBatchRequest(items, "test")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(batchWriteItemInputs))
	for i := range 3 {
		assert.Equal(t, 25, len(batchWriteItemInputs[i].RequestItems["test"]))
	}

	assert.Equal(t, 20, len(batchWriteItemInputs[3].RequestItems["test"]))
}
