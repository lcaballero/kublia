package conf

type Config struct {
	MaxPublishEntries int    `long:"max-publish-entries" description:"The size of a single partition in number of entries."`
	MaxPublishBytes   int    `long:"max-publish-bytes" description:"The size of a single partition in number of bytes."`
	BlockSize         int    `long:"block-size" description:"The block size to use as a factor for disk allocation"`
	Port              int    `long:"port" description:"Port where the webserver will bind" default:"1234"`
	Ip                string `long:"ip" description:"The IP where the webserver will bind" default:"127.0.0.1"`
	AssetRoot         string `long:"asset-root" description:"Asset root directory" default:".www/dest"`
	DefaultPage       string `long:"default-page" description:"Page to serve when matching root path" default:"index.html"`
	DeployedTo        string `long:"deployed-to" description:"The environment where the app is deployed (ie prod, dev, acc, etc)" default:"dev"`
	KeysFile          string `long:"keys-file" description:"The file from which to load secrets" default:"keys"`
	LoginPage         string `long:"login-page" description:"The page to redirect to request login or when session expires" default:"/login"`
	MimeTypes         string `long:"mime-types" description:"The mime-type file to use to provide content-types from the server" default:".files/mime-types.txt"`


	Command           string `long:"command" description:"Temporary way to run tools"`
}
