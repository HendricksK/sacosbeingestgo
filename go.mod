module github.com/HendricksK/sacosbeingestgo

// +heroku goVersion go1.15
go 1.15

require (
	github.com/HendricksK/sacosbeingestgo/ingest/githubimages v0.0.0-20210508174508-c023ccb7ed41
	github.com/HendricksK/sacosbeingestgo/ingest/sacos v0.0.0-20210508215429-cfc79eea5b7b
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/labstack/echo/v4 v4.7.2
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/HendricksK/sacosbeingestgo/ingest/sacos => ./ingest/sacos

replace github.com/HendricksK/sacosbeingestgo/ingest/githubimages => ./ingest/githubimages
