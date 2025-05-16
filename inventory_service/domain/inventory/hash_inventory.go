package inventory

import "strings"

func HashInventory(productId string, types []string) string {
	token := []string{productId}

	token = append(token, types...)

	return strings.Join(token, "-")
}
