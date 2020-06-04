package notification

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type EmailItem struct {
	Usn   string
	Name  string
	Email string
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func CreateNumericCondition(attribute string, cond string, valuestring string) expression.ConditionBuilder {
	//logger.WriteToFile("Filter Condition Being generated")

	value64, _ := strconv.ParseFloat(valuestring, 32)
	value := float32(value64)

	if cond == "Equal" {
		return expression.Name(attribute).Equal(expression.Value(value))
	}
	if cond == "GreaterThan" {
		return expression.Name(attribute).GreaterThan(expression.Value(value))
	}
	if cond == "GreaterThanEqual" {
		return expression.Name(attribute).GreaterThanEqual(expression.Value(value))
	}
	if cond == "LessThan" {
		return expression.Name(attribute).LessThan(expression.Value(value))
	}
	if cond == "LessThanEqual" {
		return expression.Name(attribute).LessThanEqual(expression.Value(value))
	}
	exitWithError(fmt.Errorf("Invalid Condtion"))
	//return empty condition
	return expression.Name("Email").AttributeExists()

}
func CreateListCondition(attribute string, values []string) expression.ConditionBuilder {
	//logger.WriteToFile("Filter Condition Being generated")
	if values[0] == "" {
		return expression.Name("Email").AttributeExists()
	}
	cond := expression.Name(attribute).Equal(expression.Value(values[0]))
	for i := 1; i < len(values); i++ {
		if values[i] != "" {
			cond = cond.Or(expression.Name(attribute).Equal(expression.Value(values[i])))
		}
	}
	return cond
}

func CreateStringCondition(attribute string, value string) expression.ConditionBuilder {
	return expression.Name(attribute).Equal(expression.Value(value))
}

func CreateBoolCondition(attribute string, value string) expression.ConditionBuilder {
	b, _ := strconv.ParseBool(value)
	return expression.Name(attribute).Equal(expression.Value(b))
}

func CreateCondition(Backlog string, StarOffer string, Branch []string, Gender string, CgpaCond string, Cgpa string, Perc10thCond string, Perc10th string, Perc12thCond string, Perc12th string) expression.ConditionBuilder {
	//logger.WriteToFile("Filter Condition Being generated")

	cond := expression.Name("Email").AttributeExists()
	//cond := expression.Name("Email").NotEqual(expression.Value(""))

	if Backlog != "" {
		cond = cond.And(CreateBoolCondition("Backlog", Backlog))
	}
	if StarOffer != "" {
		cond = cond.And(CreateBoolCondition("StarOffer", StarOffer))
	}
	if len(Branch) != 0 {
		cond = cond.And(CreateListCondition("Branch", Branch))
	}
	if Gender != "" {
		cond = cond.And(CreateStringCondition("Gender", Gender))

	}
	if CgpaCond != "" {
		cond = cond.And(CreateNumericCondition("Cgpa", CgpaCond, Cgpa))
	}
	if Perc10thCond != "" {
		cond = cond.And(CreateNumericCondition("Perc10th", Perc10thCond, Perc10th))
	}
	if Perc12thCond != "" {
		cond = cond.And(CreateNumericCondition("Perc12th", Perc12thCond, Perc12th))
	}

	return cond
}
func ApplyFilter(Backlog string, StarOffer string, Branch []string, Gender string, CgpaCond string, Cgpa string, Perc10thCond string, Perc10th string, Perc12thCond string, Perc12th string) []EmailItem {
	// create an aws session
	//logger.WriteToFile("Filter Condition Being generated")

	sess, _ := session.NewSession(&aws.Config{

		Region: aws.String("us-east-1"), DisableSSL: aws.Bool(true),
	})

	// create a dynamodb instance
	svc := dynamodb.New(sess)

	// Create the Expression to fill the input struct with.
	cond := CreateCondition(Backlog, StarOffer, Branch, Gender, CgpaCond, Cgpa, Perc10thCond, Perc10th, Perc12thCond, Perc12th)
	proj := expression.NamesList(expression.Name("Usn"), expression.Name("Name"), expression.Name("Email"))
	expr, err := expression.NewBuilder().WithCondition(cond).WithProjection(proj).Build()
	if err != nil {
		exitWithError(fmt.Errorf("failed to create the Expression, %v", err))
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Condition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Student"),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		exitWithError(fmt.Errorf("failed to make Query API call, %v", err))
	}

	emailitems := []EmailItem{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &emailitems)
	if err != nil {
		exitWithError(fmt.Errorf("failed to unmarshal Query result items, %v", err))
	}
	//fmt.Println(emailitems)
	return emailitems

}

// func main() {
// 	Backlog := "true"
// 	StarOffer := ""
// 	Branch := ""
// 	Gender := ""
// 	CgpaCond := ""
// 	//CgpaCond := "GreaterThan"
// 	Cgpa := "5"
// 	//Perc10thCond := ""
// 	Perc10thCond := "GreaterThan"
// 	Perc10th := "30"
// 	//Perc12thCond := ""
// 	Perc12thCond := "GreaterThan"
// 	Perc12th := "90"

// 	emailitems := ApplyFilter(Backlog, StarOffer, Branch, Gender, CgpaCond, Cgpa, Perc10thCond, Perc10th, Perc12thCond, Perc12th)
// 	fmt.Println(emailitems)
// }
