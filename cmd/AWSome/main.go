package main

import (
	"fmt"
	"os"

	"github.com/AWSome/pkg/policy"
)

func main() {

	d, err := os.ReadFile("/home/zuk0/Projects/AWSome/demo.json")

	if err != nil {
		fmt.Println(err)
	}

	p, err := policy.ParsePolicyDocument([]byte(d))

	if err != nil {
		fmt.Println(err)
	}

	for _, s := range p.Statement {
		fmt.Println(policy.CheckIfContainsAction("s3:ListBucket", s))
	}

}
