test:
	go test ./algo/* -v

bench:
	go test -bench=. ./algo/*

play:
	go run playground.go $(args)