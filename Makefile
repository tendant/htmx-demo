SOURCES := $(shell find . -mindepth 2 -name "main.go")
DESTS := $(patsubst ./%/main.go,dist/%,$(SOURCES))
ALL := dist/main $(DESTS)

all: $(ALL)
	@echo $@: Building Targets $^

dist/main:
ifneq (,$(wildcard main.go))
	$(echo Bulding main.go)
	go build -buildvcs -o $@ main.go
endif

#dist/main:
#	@echo Building $^ into $@
#	test -f main.go && go build -buildvcs -o $@ $^

dist/%: %/main.go
	@echo $@: Building $^ to $@
	go build -buildvcs -o $@ $^

dep:
	go mod tidy

clean:
	go clean
	rm -f $(ALL)

.PHONY: clean

css:
	npx tailwindcss -o ./static/css/output.css

migration-create:
# Usaged: make migration-create name="demo"
	migrate create -dir "migrations" -format "20060102150405" -ext sql $(name)

migrate:
	migrate -source file://migrations -database postgres://htmx_demo:pwd@localhost:5432/htmx_demo_db?sslmode=disable up

run:
	arelo -t . -p '**/*.go' -i '**/*_test.go' -i 'static/**/.*' -- go run .
