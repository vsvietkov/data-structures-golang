.PHONY: test benchmark

test:
	go test -v ./tests/...

benchmark:
	go test -v -bench=. ./benchmarks/... -benchmem
