package dynamostore

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/meokg456/productservice/dbconst"
	"github.com/meokg456/productservice/pkg/config"
)

type Options struct {
	Region   string
	Endpoint string
}

func ParseFromConfig(config *config.Config) Options {
	return Options{
		Region:   config.DB.Region,
		Endpoint: config.DB.Endpoint,
	}
}

func NewConnection(options Options) (*dynamodb.Client, error) {
	cfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(options.Region),
		awsconfig.WithBaseEndpoint(options.Endpoint),
	)

	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return client, nil
}

func CheckMigratedTable(db *dynamodb.Client, migratedTables map[string]bool) {
	input := &dynamodb.ListTablesInput{}
	result, err := db.ListTables(context.Background(), input)
	if err != nil {
		log.Fatalln(err)
	}

	for _, tableName := range result.TableNames {
		migratedTables[tableName] = true
	}
}

func CreateTable(db *dynamodb.Client, tableName string) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeType("S"),
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyType("HASH"),
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := db.CreateTable(context.Background(), input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	log.Printf("Created table %s", tableName)
}

func MigrateDatabase(db *dynamodb.Client) {
	tables := []string{dbconst.ProductTableName}

	migratedTables := map[string]bool{}

	for _, table := range tables {
		migratedTables[table] = false
	}

	CheckMigratedTable(db, migratedTables)

	for _, table := range tables {
		migrated := migratedTables[table]
		if !migrated {
			CreateTable(db, table)
		}
	}
}
