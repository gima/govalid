## Contributing

Use pull requests or file an issue and relay your intent for the action.

Try to avoid mixing different concerns in one commit. *(look who's talking)* Same applies to pull requests and issues.

### Code formatting

Run `go fmt` before committing.

### Run tests

Test all reasonable code paths.

    go test ./tests/

### Run code coverage

Try to cover the necessary cases.

    go test -cover ./tests/ -coverpkg ./ -coverprofile cover.out; go tool cover -html=cover.out -o coverage.html

*This contribution readme was shamelessly modelled after:  
http://opencomparison.readthedocs.org/en/latest/contributing.html*
