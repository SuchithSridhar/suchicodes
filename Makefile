VIEWS_DIR = ./web/views/
SCSS_SRC_DIR = ./web/scss/
CSS_OUTPUT_DIR = ./web/static/css/
NOTIFICATION_TITLE = "Suchicodes Makefile"

run: gen_templ gen_sass
	@go run main.go

gen_templ:
	# generate templates in 
	@templ generate -path $(VIEWS_DIR)

gen_sass:
	@mkdir -p $(CSS_OUTPUT_DIR)
	@sass $(SCSS_SRC_DIR):$(CSS_OUTPUT_DIR)

watch_sass:
	@mkdir -p $(CSS_OUTPUT_DIR)
	@sass --watch $(SCSS_SRC_DIR):$(CSS_OUTPUT_DIR)

fmt:
	# format all Go source files
	@find . -name '*.go' -exec gofmt -s -w {} \+
	@templ fmt $(VIEWS_DIR)
	@find . -iname "*.json" -exec bash -c "jq . {} > /dev/null && jq . {} | sponge {} " \;
	@notify-send -t 5000 -u normal -a $(NOTIFICATION_TITLE) "Formatted Go and Templ files" 

kill:
	# kill running Go server
	@-pkill -f "go run main.go"

.PHONY: run gen_templ gen_sass fmt kill

