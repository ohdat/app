.PHONY: check
check:
	go fmt ./*
	go vet ./*