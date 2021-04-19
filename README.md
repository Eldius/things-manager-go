# things-manager-go #

## snippets ##

```shell
# add test thing 0
curl -i -XPOST localhost:8080/things -d '{
    "name": "Test Thing",
    "description": "A simple test thing...",
    "available": 0
}'

# add test thing 1
curl -i -XPOST localhost:8080/things -d '{
    "name": "Test Thing 01",
    "description": "A simple test thing 02...",
    "available": 0
}'

# list things
curl -i -XGET localhost:8080/things
```
