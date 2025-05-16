package inventory

import "strings"

func HashInventory(inventory Inventory) string {
	token := []string{inventory.ProductId}

	token = append(token, inventory.Types...)

	return strings.Join(token, "-")
}
