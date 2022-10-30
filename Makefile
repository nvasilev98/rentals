.PHONY: unit-tests

unit-tests:
	ginkgo -r -race -randomize-all -randomize-suites .