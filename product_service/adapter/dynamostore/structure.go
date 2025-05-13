package dynamostore

type ProductData struct {
	ID   string `dynamodbav:"ID"`
	Name string `dynamodbav:"Name"`
}
