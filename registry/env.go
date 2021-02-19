package registry

type Envs struct {
	KMSClient *adapter.AWSKmsClient
	Cache map[string]string
}

var envs *Envs

func NewEnvs() *Envs {
	return &Envs{
		KMSClient: adapter.NewAWSKmsClient()k,
		Cache: make(map[string]string),
	}
}

func Env() *Envs {
	if envs == nil {
		envs = NewEnvs()
	}
}

func (c *Envs) env(key string) string {
	return os.Getenv(key)
}

func (c *Envs) DynamoLocalEndpoint() string {
	return c.env("DDYNAMO_LOCAL_ENDPOINT")
}

func (c *Envs) DynamoTableName() string {
	return c.env("DYNAMO_TABLE_NAME")
}

func (c *Envs) DynamoPKName() string {
	return c.env("DYNAMO_PK_NAME")
}

func (c *Envs) DynamoSKName() string {
	return c.env("DYNAMO_SK_NAME")
}

