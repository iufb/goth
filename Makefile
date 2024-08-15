run: build
	@./bin/app

build:
	@go build -o bin/app .


css:
	tailwindcss -i views/globals.css -o public/styles.css --watch   
