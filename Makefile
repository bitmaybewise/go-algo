test:
	go test ./algo/*

test-verbose:
	go test ./algo/* -v

test1:
	go test -v ./algo/* -run $(name)

bench:
	go test -bench=. ./algo/*

play:
	go run playground.go $(args)
