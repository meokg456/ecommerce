package dynamostore

type ProductData struct {
	ID           string         `dynamodbav:"ID"`
	Title        string         `dynamodbav:"Title"`
	Descriptions string         `dynamodbav:"Descriptions"`
	Category     string         `dynamodbav:"Category"`
	Images       []string       `dynamodbav:"Images"`
	AdditionInfo map[string]any `dynamodbav:"AdditionInfo"`
	MerchantId   int            `dynamodbav:"MerchantId"`
}
