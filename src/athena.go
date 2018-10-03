package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func main() {
	awscfg := &aws.Config{}
	awscfg.WithRegion("ap-southeast-2")
	// Create the session that the service will use.
	sess := session.Must(session.NewSession(awscfg))
	svc := athena.New(sess, aws.NewConfig().WithRegion("ap-southeast-2"))
	var s athena.StartQueryExecutionInput
	s.SetQueryString("select PageURL from testtable limit 10")

	var q athena.QueryExecutionContext
	q.SetDatabase("testdb")
	s.SetQueryExecutionContext(&q)

	var r athena.ResultConfiguration
	r.SetOutputLocation("s3://datalake-gez")
	s.SetResultConfiguration(&r)

	result, err := svc.StartQueryExecution(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(result.GoString())

}
