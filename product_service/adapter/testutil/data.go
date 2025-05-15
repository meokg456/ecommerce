package testutil

import "github.com/meokg456/productservice/domain/product"

var Products = []product.Product{
	{
		Id:           "1",
		Title:        "Laptop",
		Descriptions: "A high-performance laptop for work and gaming.",
		Category:     "Electronics",
		Images:       []string{"laptop1.jpg", "laptop2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandX", "warranty": "2 years"},
		MerchantId:   1,
	},
	{
		Id:           "2",
		Title:        "Smartphone",
		Descriptions: "A sleek smartphone with a powerful camera.",
		Category:     "Electronics",
		Images:       []string{"smartphone1.jpg", "smartphone2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandY", "battery": "4000mAh"},
		MerchantId:   1,
	},
	{
		Id:           "3",
		Title:        "Headphones",
		Descriptions: "Noise-cancelling over-ear headphones.",
		Category:     "Accessories",
		Images:       []string{"headphones1.jpg", "headphones2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandZ", "type": "Wireless"},
		MerchantId:   1,
	},
	{
		Id:           "4",
		Title:        "Keyboard",
		Descriptions: "Mechanical keyboard with RGB lighting.",
		Category:     "Accessories",
		Images:       []string{"keyboard1.jpg", "keyboard2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandA", "switchType": "Cherry MX Red"},
		MerchantId:   1,
	},
	{
		Id:           "5",
		Title:        "Smartwatch",
		Descriptions: "A smartwatch with fitness tracking features.",
		Category:     "Wearables",
		Images:       []string{"smartwatch1.jpg", "smartwatch2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandB", "waterproof": true},
		MerchantId:   1,
	},
}
