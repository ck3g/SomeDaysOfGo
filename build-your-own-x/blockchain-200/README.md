# Blockchain

[Code your own blockchain in less than 200 lines of Go!](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)

## How to use

1. Start the server `go run main.go`
2. Create new block `curl -d '{"BPM":60}' -H 'Content-Type: application/json' http://localhost:8080` (change `60` to a different number)
3. Visit `http://localhost:8080` (or use `curl http://localhost:8080`) to see the results