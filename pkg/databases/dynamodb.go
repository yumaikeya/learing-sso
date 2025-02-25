package databases

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoのLocal
func NewLocalDynamo() *dynamodb.Client {
	// customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	return aws.Endpoint{
	// 		PartitionID:   "aws",
	// 		URL:           "http://db:8000/",
	// 		SigningRegion: "localhost:8000",
	// 	}, nil
	// })

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("a", "b", "c")),
		// config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		fmt.Printf("Fail to create AWS config: %#v", err.Error())
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://db:8000/")
	})
}

// DEV, PROD
func NewDynamo() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"), config.WithSharedConfigProfile("MAPPER-Dev"))
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func InjectNextToken(input *dynamodb.QueryInput, pk, sk string) {
	input.ExclusiveStartKey = map[string]types.AttributeValue{
		"pk": &types.AttributeValueMemberS{Value: pk},
		"sk": &types.AttributeValueMemberS{Value: sk},
	}
}

func InjectNextTokenScan(input *dynamodb.ScanInput, pk, sk types.AttributeValue) {
	input.ExclusiveStartKey = map[string]types.AttributeValue{
		"pk": pk,
		"sk": sk,
	}
}

// equalによってscanしたい時
func MakeScanInputEqual(tableName, sk string) *dynamodb.ScanInput {
	expr, _ := expression.NewBuilder().WithFilter(expression.Equal(expression.Name("sk"), expression.Value(sk))).Build()

	return &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}
}

// containによってscanしたい時
func MakeScanInputContain(tableName, kinda string, sortBy *string, limit *int32, opts ...*[]string) *dynamodb.ScanInput {
	filterCond := expression.Contains(expression.Name("sk"), kinda)
	for h := range opts {
		if opts[h] == nil {
			continue
		}

		filterCond = filterCond.And(func() (kwdFilterCond expression.ConditionBuilder) {
			for i := range *opts[h] {
				if i == 0 {
					kwdFilterCond = expression.Contains(expression.Name("name"), (*opts[h])[i]).Or(expression.Contains(expression.Name("organizationName"), (*opts[h])[i])).Or(expression.Contains(expression.Name("tagIds"), (*opts[h])[i]))
				}
				kwdFilterCond = kwdFilterCond.Or(expression.Contains(expression.Name("name"), (*opts[h])[i]), expression.Contains(expression.Name("organizationName"), (*opts[h])[i]), expression.Contains(expression.Name("tagIds"), (*opts[h])[i]))
			}
			return kwdFilterCond
		}()) //TODO: projectの検索用...ここにあるのは不本意ではある
	}
	expr, _ := expression.NewBuilder().WithFilter(filterCond).Build()

	return &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		Limit:                     limit,
	}
}

// beginWithによってqueryしたい時
func MakeQueryInputBegin(tableName, pk, sk string, order *bool, limit *int32, opts ...*[]string) *dynamodb.QueryInput {
	keyCond := expression.Key("project_id").Equal(expression.Value(pk)).And(expression.KeyBeginsWith(expression.Key("sk"), sk))
	filterCond := expression.NotEqual(expression.Name("id"), expression.Value(""))

	expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(filterCond).Build()

	return &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          order,
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
		Limit:                     limit,
	}
}

// ids検索用
func MakeQueryInputIds(tableName, pk string, ids []string) *dynamodb.QueryInput {
	filterCond := func() (kwdFilterCond expression.ConditionBuilder) {
		for i := range ids {
			if i == 0 {
				kwdFilterCond = expression.Contains(expression.Name("id"), ids[i])
			}
			kwdFilterCond = kwdFilterCond.Or(expression.Contains(expression.Name("id"), ids[i]))
		}
		return kwdFilterCond
	}()

	expr, _ := expression.NewBuilder().WithKeyCondition(expression.Key("project_id").Equal(expression.Value(pk))).WithFilter(filterCond).Build()

	return &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
	}
}

// 現在時間でactiveなqueryを生成したい時
func MakeQueryInputActive(tableName, pk string, isActive bool, limit *int32) *dynamodb.QueryInput {
	now := time.Now()
	keyCond := expression.Key("project_id").Equal(expression.Value(pk))
	filterCond := expression.NotEqual(expression.Name("activatesIn"), expression.Value(""))
	if isActive {
		keyCond = keyCond.And(expression.KeyGreaterThan(expression.Key("expiresIn"), expression.Value(now)))
		filterCond = expression.LessThan(expression.Name("activatesIn"), expression.Value(now))
	}

	expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(filterCond).Build()

	return &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
		Limit:                     limit,
	}
}

// idxを使ってqueryしたい時
func MakeQueryInputIdx(tableName, idxName, idxValue string) *dynamodb.QueryInput {
	keyEx := expression.Key(idxName).Equal(expression.Value(idxValue))
	expr, _ := expression.NewBuilder().WithKeyCondition(keyEx).Build()

	return &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String(fmt.Sprintf("gsi-%s", idxName)),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}
}

// lsi-idxを使ってquery(get)したい時
func MakeQueryInputLsiIdx(tableName, projectId, idxName, idxValue string) *dynamodb.QueryInput {
	keyEx := expression.Key("project_id").Equal(expression.Value(projectId)).And(expression.KeyBeginsWith(expression.Key(idxName), idxValue))
	expr, _ := expression.NewBuilder().WithKeyCondition(keyEx).Build()

	return &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String(fmt.Sprintf("lsi-%s", idxName)),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}
}

// DeleteItem作成用
func MakeDeleteInput(tableName, pk, sk string) *dynamodb.DeleteItemInput {
	return &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"project_id": &types.AttributeValueMemberS{Value: pk},
			"sk":         &types.AttributeValueMemberS{Value: sk},
		},
	}
}

// DeleteItem作成用
func MakeBulkDeleteInput(tableName string, pksks []map[string]string) *dynamodb.TransactWriteItemsInput {
	txItems := []types.TransactWriteItem{}
	for _, pksk := range pksks {
		txItems = append(txItems, types.TransactWriteItem{
			Delete: &types.Delete{
				TableName: aws.String(tableName),
				Key: map[string]types.AttributeValue{
					"project_id": &types.AttributeValueMemberS{Value: pksk["project_id"]},
					"sk":         &types.AttributeValueMemberS{Value: pksk["sk"]},
				},
			},
		})
	}

	return &dynamodb.TransactWriteItemsInput{
		TransactItems: txItems,
	}
}

func MakeUpdateInput(tableName, pk, sk string, data any) (*dynamodb.UpdateItemInput, error) {
	updateValues, err := attributevalue.MarshalMap(data)
	if err != nil {
		return nil, err
	}

	formattedUpdateValues := make(map[string]types.AttributeValue)
	for k, v := range updateValues {
		formattedUpdateValues[":"+k] = v
	}

	//これらがあるとエラる
	delete(formattedUpdateValues, ":project_id")
	delete(formattedUpdateValues, ":sk")

	updateNames, updateString := makeUpdateExpression(formattedUpdateValues)

	return &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"project_id": &types.AttributeValueMemberS{Value: pk},
			"sk":         &types.AttributeValueMemberS{Value: sk},
		},
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  updateNames,
		ExpressionAttributeValues: formattedUpdateValues,
		UpdateExpression:          updateString,
		ReturnValues:              types.ReturnValueAllNew,
	}, nil
}

func makeUpdateExpression(attributeValue map[string]types.AttributeValue) (map[string]string, *string) {
	updateNames := make(map[string]string, len(attributeValue))
	updateExpression := "SET "
	for k := range attributeValue {
		key1 := strings.Replace(k, ":", "", 1)
		key2 := "#" + key1
		updateNames[key2] = key1                                                 //"#coordinates": aws.String("coordinates")
		updateExpression = updateExpression + fmt.Sprintf("%v=:%v,", key2, key1) //SET #coordinates=:coordinates
	}
	updateExpression = updateExpression[:len(updateExpression)-1]
	return updateNames, aws.String(updateExpression)
}
