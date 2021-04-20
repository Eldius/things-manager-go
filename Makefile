
start:
	THINGS_APP_DATABASE_LOG=1 go run main.go start

addtestthings:
	# add test thing 0
	curl -i -XPOST localhost:8080/things -d '{"name": "Test Thing", "description": "A simple test thing...", "available": 0}'
	# add test thing 1
	curl -i -XPOST localhost:8080/things -d '{"name": "Test Thing 01", "description": "A simple test thing 01...", "available": 0}'
