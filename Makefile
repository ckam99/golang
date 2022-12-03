
gen:
	rm -rf pb/*.proto
	cd proto;buf generate
update:
	cd proto;buf mod update
serve:
	go run cmd/main.go
evans:
	evans --host localhost --port 6000 -r repl