module github.com/HendricksK/sacosbeingestgo

// +heroku goVersion go1.15
go 1.15

require (
	github.com/HendricksK/sacosbeingestgo/ingest/githubimages v0.0.0-20210508174508-c023ccb7ed41 // indirect
	github.com/HendricksK/sacosbeingestgo/ingest/sacos v0.0.0-20210508215429-cfc79eea5b7b
	github.com/canthefason/go-watcher v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.7.2
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/HendricksK/sacosbeingestgo/ingest/sacos => ./ingest/sacos

replace github.com/HendricksK/sacosbeingestgo/ingest/githubimages => ./ingest/githubimages
