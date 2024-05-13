VIEWS_DIR = ./web/views/
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
	@notify-send -t 5000 -u normal -a $(NOTIFICATION_TITLE) "Formatted Go and Templ files" 

kill:
	# kill running Go server
	@-pkill -f "go run main.go"

.PHONY: run gen fmt kill

