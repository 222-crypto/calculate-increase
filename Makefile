SHELL := /bin/bash
APP_NAME := $(shell basename $(shell pwd))
# APP_NAME := $(shell git remote get-url origin | awk '{split($$0,a,"/");print a[2]}' | sed 's/\.git//g')

.PHONY: help # 🙂
help: .gitignore bin/
	@printf "Available targets for \033[92m$(APP_NAME)\033[0m:\n\n"
	@cat Makefile  | grep ".PHONY" | grep -v ".PHONY: _" | awk '{split($$0,a,".PHONY: ");split(a[2],b,"#");print "\033[36m"b[1]"\033[0m#"b[2]}' | column -s '#' -t


.PHONY: build # 🙂
build: bin/$(APP_NAME) bin/verse
bin/$(APP_NAME): go.mod calculated.go
	go build -o bin/$(APP_NAME) .


.PHONY: run # 🤔
run: build calculated.csv
	bat calculated.csv
	@./bin/verse | cowsay -f ./share/cows/golgotha.cow | lolcat


.gitignore:
	printf "/bin/\n" > .gitignore


bin:
	mkdir -p bin


go.mod:
	go mod init $(APP_NAME)


deb/verse_0.22.9.tar.xz:
	mkdir -p deb
	curl -s "https://cdn-aws.deb.debian.org/debian/pool/main/v/verse/verse_0.22.9.tar.xz" -o ./deb/verse_0.22.9.tar.xz
	# Confirm Expectation
	diff \
		<(sha512sum deb/verse_0.22.9.tar.xz) \
		<(echo "0843664842b5f18812c0b9bf2542169fe2ddbaf6a7852d961d0d776e541a988da5ed5d3ea5afdb996a81762ce1e3e8d7c62143dc2b5cc26780c93430a35844ce  deb/verse_0.22.9.tar.xz") \
	;


src/verse-0.22.9/verse.c:
	$(MAKE) deb/verse_0.22.9.tar.xz
	mkdir -p src
	tar -xvf ./deb/verse_0.22.9.tar.xz -C ./src


bin/verse: bin src/verse-0.22.9/verse.c
	cd src/verse-0.22.9 && make verse VERSE_LIB=$$(pwd)/ && mv verse ../../bin


calculated.csv: bin/$(APP_NAME)
	bin/$(APP_NAME) > calculated.csv
	bat calculated.csv
