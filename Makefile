

gen:
	chmod +x ./bin/tools/zek
	mkdir -p pkg/openimmo
	./bin/tools/zek -j -C < ./xds/v1.2.7b_16_06_2020/openimmo-data_127.xml > ./pkg/openimmo/model.go


tools: 	./bin/tools/zek

./bin/tools/zek:
	mkdir -p ./bin/tools
	curl -L -o ./bin/tools/zek https://github.com/miku/zek/releases/download/v0.1.9/zek-0.1.9.darwin.amd64.bin
