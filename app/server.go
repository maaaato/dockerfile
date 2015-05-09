package main


import (
	"os"
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
	
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/ec2"	
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {

	access_key := os.Getenv("MAAAATO_AWS_ACCESS_KEY_ID")
	secret_key := os.Getenv("MAAAATO_AWS_SECRET_ACCESS_KEY")
	
	creds := aws.Creds(access_key, secret_key, "")
	cli := ec2.New(creds, "ap-northeast-1", nil)
	resp, _ := cli.DescribeInstances(nil)
	fmt.Fprintf(w, *resp.Reservations[0].Instances[0].VPCID)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Serve()
}
