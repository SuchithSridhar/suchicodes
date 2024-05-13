VIEWS_DIR = ./web/views/
NOTIFICATION_TITLE = "Suchicodes Makefile"

run: gen
	@go run main.go &

gen:
	# generate templates in 
	@templ generate -path $(VIEWS_DIR)

fmt:
	# format all Go source files
	@find . -name '*.go' -exec gofmt -s -w {} \+
	@echo "Formatted Go files" | notify-send -t 5000 -u normal $(NOTIFICATION_TITLE)
	@templ fmt $(VIEWS_DIR)
	@echo "Formatted Templ files" | notify-send -t 5000 -u normal $(NOTIFICATION_TITLE)

watch:
	# watch for changes in any files and recompile using inotifywait
	@while inotifywait -r -e modify,create,delete --exclude '\.git|node_modules|\.swp' .; do \
		make kill gen run; \
	done

kill:
	# kill running Go server
	@-pkill -f "go run main.go"

all: fmt fmt-templ run
	@echo "Recompiled the project" | notify-send -t 5000 -u normal $(NOTIFICATION_TITLE)

.PHONY: run gen fmt watch all kill

