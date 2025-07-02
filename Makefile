.PHONY: test test-dev badge

test:
	go test -coverprofile=coverage.out

test-dev:
	go test -tags dev -coverprofile=coverage-dev.out

badge: test-dev
	go run tools/genbadge.go
