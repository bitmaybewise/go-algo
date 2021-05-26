test:
	go test ./...

test-verbose:
	go test ./... -v

test1:
	go test -v ./... -run $(name)

bench:
	go test -bench=. ./...

play:
	go run playground.go $(args)
