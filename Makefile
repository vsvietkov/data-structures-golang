.PHONY: test benchmark

test:
	go test ./tests/...

benchmark:
	go test -v -bench=. ./benchmarks/... -benchmem
