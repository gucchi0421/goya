.PHONY: i
i:
	@go get -u ./...
	@go install github.com/spf13/cobra-cli@latest


.PHONY: run
run:
	go run . $(filter-out $@,$(MAKECMDGOALS))

%:
	@: