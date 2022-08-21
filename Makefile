# regenerate parser
.PHONY: generate_parser
generate_parser:
	@java -jar /usr/local/lib/antlr-4.10.1-complete.jar -Dlanguage=Go ./parser/JsonPath.g4

.PHONY: test
test:
	@go test -v ./...	
