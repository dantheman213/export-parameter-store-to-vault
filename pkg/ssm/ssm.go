package ssm

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ssm"
)

var s *session.Session
var svc *ssm.SSM

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
func InitSsm(profile string) {
    s = session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        Profile: profile,
    }))
    svc = ssm.New(s)
}

func ListAllParameters() ([]string, error) {
    response := make([]string, 0)

    maxAllowedResults := int64(50)
    temp := ""
    var next = &temp

    for next != nil {
        if *next == "" {
            next = nil
        }
        output, err := svc.DescribeParameters(&ssm.DescribeParametersInput{
            Filters:          nil,
            MaxResults:       &maxAllowedResults,
            NextToken:        next,
            ParameterFilters: nil,
        })
        if err != nil {
            return nil, err
        }

        for _, p := range output.Parameters {
            response = append(response, *p.Name)
        }
        next = output.NextToken
    }

    return response, nil
}

func ReadParameterValue(key string) (string, error) {
    output, err := svc.GetParameter(&ssm.GetParameterInput{
        Name:           &key,
        WithDecryption: nil,
    })
    if err != nil {
        return "", err
    }

    return *output.Parameter.Value, nil
}