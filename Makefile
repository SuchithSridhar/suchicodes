VIEWS_DIR = ./web/
NOTIFICATION_TITLE = "Suchicodes Makefile"

run: gen
	@go run main.go

gen:
	# generate templates in 
	@templ generate -path $(VIEWS_DIR)

fmt:
	# format all Go source files
	@find . -name '*.go' -exec gofmt -s -w {} \+
	@templ fmt $(VIEWS_DIR)
	@find . -iname "*.json" -exec bash -c "jq . {} > /dev/null && jq . {} | sponge {} " \;
	@notify-send -t 5000 -u normal -a $(NOTIFICATION_TITLE) "Formatted Go and Templ files" 

kill:
	# kill running Go server
	@-pkill -f "go run main.go"

.PHONY: run gen fmt kill

