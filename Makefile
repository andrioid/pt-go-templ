BINARY_NAME=main

build: $(wildcard *.go) $(wildcard *.templ) tailwind
	go mod tidy
	templ generate
	go build -o ${BINARY_NAME} ./

clean:
	go clean
	rm ${BINARY_NAME}

tailwind:
	npx -y tailwindcss -i ./web/global.css -o ./web/public/generated.css

run: tailwind build
	./main

watch:
	air -c .air.toml