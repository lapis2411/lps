cp_template:
	cp ./config/config.template.yml ./config/config.yml
build:
	go build -ldflags="-s -w" -trimpath