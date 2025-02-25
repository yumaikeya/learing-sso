package main

import (
	"angya-backend/pkg/constants"
	"angya-backend/pkg/databases"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// 初期時にDynamoDBにtablesを作成するScript
func main() {
	ctx, client := context.TODO(), databases.NewLocalDynamo() // DynamoLocal
	res, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		print("Err: ", err.Error())
		panic("Can't list tables")
	}

	if len(res.TableNames) != 0 {
		for i := range res.TableNames {
			_, err := client.DeleteTable(ctx, &dynamodb.DeleteTableInput{
				TableName: aws.String(res.TableNames[i]),
			})
			if err != nil {
				panic("Can't delete a table which already been...")
			}
		}
	}

	if _, err = client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(constants.PLANNER_TABLE),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("project_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("project_id"),
				KeyType:       types.KeyTypeHash, // HASH(Partition key)を設定
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       types.KeyTypeRange, // RANGE(Sort key)を設定
			},
		},
		LocalSecondaryIndexes: []types.LocalSecondaryIndex{
			{
				IndexName: aws.String("lsi-id"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("project_id"),
						KeyType:       types.KeyTypeHash, // HASH(Partition key)を設定
					},
					{
						AttributeName: aws.String("id"),
						KeyType:       types.KeyTypeRange, // Range(Sort key)を設定
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}); err != nil {
		fmt.Printf("Err: %#v", err.Error())
		panic("Can't create project table")
	}
}

// type flight struct {
// 	ProjectId string  `dynamodbav:"pk"`
// 	Sk        string  `dynamodbav:"sk"`
// 	PointSize float32 `dynamodbav:"pointSize"`
// }

// // poiのtypeを埋めるscript
// func main() {
// 	ctx, ddb := context.TODO(), databases.NewDynamo()

// 	filterCond := expression.Contains(expression.Name("sk"), "flight-").And(expression.AttributeNotExists(expression.Name("pointSize")))
// 	expr, _ := expression.NewBuilder().WithFilter(filterCond).Build()

// 	input, dbsAll := &dynamodb.ScanInput{
// 		TableName:                 aws.String(constants.PROJECT_TABLE),
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		FilterExpression:          expr.Filter(),
// 	}, []flight{}

// 	for {
// 		out, err := ddb.Scan(ctx, input)
// 		if err != nil {
// 			print(err.Error())
// 			panic("Failed!!!")
// 		}

// 		var dbs []flight
// 		attributevalue.UnmarshalListOfMaps(out.Items, &dbs)
// 		dbsAll = append(dbsAll, dbs...)

// 		if out.LastEvaluatedKey != nil {
// 			input.ExclusiveStartKey = out.LastEvaluatedKey
// 		} else {
// 			break
// 		}
// 	}

// 	for _, f := range dbsAll {
// 		fmt.Printf("%s  %s\n", f.ProjectId, f.Sk)

// 		update := expression.Set(expression.Name("pointSize"), expression.Value(2))
// 		expr, _ := expression.NewBuilder().WithUpdate(update).Build()

// 		_, err := ddb.UpdateItem(ctx, &dynamodb.UpdateItemInput{
// 			TableName: aws.String(constants.PROJECT_TABLE),
// 			Key: map[string]types.AttributeValue{
// 				"pk": &types.AttributeValueMemberS{Value: f.ProjectId},
// 				"sk": &types.AttributeValueMemberS{Value: f.Sk},
// 			},
// 			ExpressionAttributeNames:  expr.Names(),
// 			ExpressionAttributeValues: expr.Values(),
// 			UpdateExpression:          expr.Update(),
// 		})
// 		if err != nil {
// 			print(err.Error())
// 			panic("Failed!!!")
// 		}
// 	}
// }
