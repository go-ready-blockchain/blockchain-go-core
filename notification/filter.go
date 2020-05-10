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

func CreateCondition(Backlog string, StarOffer string, Branch string, Gender string, CgpaCond string, Cgpa string, Perc10thCond string, Perc10th string, Perc12thCond string, Perc12th string) expression.ConditionBuilder {

	cond := expression.Name("Email").AttributeExists()
	//cond := expression.Name("Email").NotEqual(expression.Value(""))

	if Backlog != "" {
		b, _ := strconv.ParseBool(Backlog)
		cond = cond.And(expression.Name("Backlog").Equal(expression.Value(b)))
	}
	if StarOffer != "" {
		b, _ := strconv.ParseBool(StarOffer)
		cond = cond.And(expression.Name("StarOffer").Equal(expression.Value(b)))
	}
	if Branch != "" {
		cond = cond.And(expression.Name("Branch").Equal(expression.Value(Branch)))
	}
	if Gender != "" {
		cond = cond.And(expression.Name("Gender").Equal(expression.Value(Gender)))
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
func ApplyFilter(Backlog string, StarOffer string, Branch string, Gender string, CgpaCond string, Cgpa string, Perc10thCond string, Perc10th string, Perc12thCond string, Perc12th string) []EmailItem {
	// create an aws session
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
