build:
	go build -ldflags '-s -w' -trimpath ./cmd/go-grep/

clean:
	rm go-grep