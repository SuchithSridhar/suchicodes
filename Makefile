run: gen
	@go run main.go

gen:
	# generate templates in view/ directory
	@templ generate
