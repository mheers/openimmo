gen:
	mkdir -p pkg/openimmo
	zek -j -C < ./xds/v1.2.7b_16_06_2020/openimmo-data_127.xml > ./pkg/openimmo/model.go

tools: 	install-zek

install-zek:
	go install github.com/miku/zek/cmd/zek@latest
