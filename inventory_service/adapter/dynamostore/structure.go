package dynamostore

type InventoryData struct {
	ID       string `dynamodbav:"ID"`
	Quantity int    `dynamodbav:"Quantity"`
}
