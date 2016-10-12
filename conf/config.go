package conf

type Config struct {
	MaxPublishSize int `long:"max-publish-size" description:"The size of a single partition in number of entries not bytes."`
}
