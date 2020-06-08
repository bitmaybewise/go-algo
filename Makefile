test:
	go test ./algo/* -v

test1:
	go test ./algo/* -run $(name)

bench:
	go test -bench=. ./algo/*

play:
	go run playground.go $(args)