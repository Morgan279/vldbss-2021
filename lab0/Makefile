.PHONY: all

all: test_example test_homework cleanup gendata

test_example:
	go test -v -timeout 60m -run=TestExampleURLTop

test_homework:
	go test -v -timeout 60m -run=TestURLTop

cleanup:
	go test -v -run=TestCleanData

gendata:
	go test -v -run=TestGenData
