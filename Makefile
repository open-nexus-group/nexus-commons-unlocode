.PHONY: generate-subdiv
generate-subdiv:
	go run ./generator/generator.go --type=subdivision --files=./generator/2024-2-SubDiv.csv

.PHONY: generate-unlocode
generate-unlocode:
	go run ./generator/generator.go --type=unlocode --files=./generator/2024-2-UNLOC-1.csv,./generator/2024-2-UNLOC-2.csv,./generator/2024-2-UNLOC-3.csv
