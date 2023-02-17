all: compiler simulator

compiler:
	cd assembler && go install

simulator:
	cd machine && go install

.PHONY: all