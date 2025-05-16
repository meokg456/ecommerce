package testutil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BuildSuccessBody(t testing.TB, data any, status int) string {
	t.Helper()

	jsonString, err := json.Marshal(map[string]any{
		"code":    strconv.Itoa(status),
		"message": "OK",
		"result":  data,
	})
	assert.NoError(t, err)

	return string(jsonString) + "\n"
}

func BuildSuccessBodyWithPagination(t testing.TB, status int, data any, page int, limit int, total int) string {
	t.Helper()

	jsonString, err := json.Marshal(map[string]any{
		"code":    strconv.Itoa(status),
		"message": "OK",
		"result": map[string]any{
			"data":  data,
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
	assert.NoError(t, err)

	return string(jsonString) + "\n"
}

func BuildErrorBody(t testing.TB, status int, errCode int) string {
	t.Helper()

	messages := "Error!"
	if status == http.StatusBadRequest {
		messages = http.StatusText(status)
	}

	jsonString, err := json.Marshal(map[string]any{
		"code":    strconv.Itoa(status) + fmt.Sprintf("%03d", errCode),
		"message": messages,
		"info":    messages,
	})
	assert.NoError(t, err)

	return string(jsonString) + "\n"
}
