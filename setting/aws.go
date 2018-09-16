package setting

type AwsProfileConfig struct {
	Profiles map[string]AwsProfile
}

type AwsProfile struct {
	Output AwsCliOutput `ini:"output"`
	Region AwsRegion    `ini:"region"`
}

type AwsCredential struct {
	AccessKeyId     string `ini:"aws_access_key_id"`
	SecretAccessKey string `ini:"aws_secret_access_key"`
}

// aws region type
// see https://docs.aws.amazon.com/general/latest/gr/rande.html
type AwsRegion string

const (
	AwsRegionUsEastVirginia   AwsRegion = "us-east-1"
	AwsRegionUsEastOhio       AwsRegion = "us-east-2"
	AwsRegionUsWestCalifornia AwsRegion = "us-west-1"
	AwsRegionUsWestOregon     AwsRegion = "us-west-2"
	AwsRegionAsiaSydney       AwsRegion = "ap-southeast-1"
	AwsRegionAsiaMumbai       AwsRegion = "ap-southeast-2"
	AwsRegionAsiaTokyo        AwsRegion = "ap-northeast-1"
	AwsRegionAsiaSeoul        AwsRegion = "ap-northeast-2"
	AwsRegionAsiaOsaka        AwsRegion = "ap-northeast-3"
	AwsRegionAsiaSingapore    AwsRegion = "ap-south-1"
	AwsRegionCanadaCentral    AwsRegion = "ca-central-1"
	AwsRegionChinaBeijing     AwsRegion = "cn-north-1"
	AwsRegionChinaNingxia     AwsRegion = "cn-northwest-1"
	AwsRegionEuFrankfurt      AwsRegion = "eu-central-1"
	AwsRegionEuIreland        AwsRegion = "eu-west-1"
	AwsRegionEuLondon         AwsRegion = "eu-west-2"
	AwsRegionEuParis          AwsRegion = "eu-west-3"
	AwsRegionEuSaoPaulo       AwsRegion = "sa-east-1"
)

// aws cli output format type
// see https://docs.aws.amazon.com/cli/latest/userguide/controlling-output.html
type AwsCliOutput string

const (
	AwsCliOutputJson  AwsCliOutput = "json"
	AwsCliOutputText  AwsCliOutput = "text"
	AwsCliOutputTable AwsCliOutput = "table"
)
