# IPHarbor

IPHarbor is a service that connects to various AI providers to generate content based on user prompts.

## How to Run IPHarbor Services

### 1. Start the JavaScript Services

```bash
cd /c/Users/vaidi/OneDrive/Desktop/IPHarbor/services
./start.sh
```

This will start the JavaScript services on port 3000.

### 2. Start the Go Backend

```bash
cd /c/Users/vaidi/OneDrive/Desktop/IPHarbor/agents
go run main.go
```

This will start the Go backend on port 8080.

### 3. Test the Services

You can use curl or any API client to test the services:

#### Generate an Image

```bash
curl -X POST http://localhost:8080/generateTextToImage \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "Astronaut riding a horse",
    "provider": "huggingface",
    "model": "black-forest-labs/FLUX.1-dev",
    "token": "hf_xxxxxxxxxxxxxxxxxxxxxxxx"
  }' \
  --output image.png
```

#### Generate Text

```bash
curl -X POST http://localhost:8080/generateTextToText \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "Once upon a time",
    "provider": "huggingface",
    "model": "gpt2",
    "token": "hf_xxxxxxxxxxxxxxxxxxxxxxxx"
  }'
```

## Service Architecture

1. Client sends a request to the Go backend
2. Go backend forwards the request to the JavaScript services
3. JavaScript services analyze the request and determine which API to use
4. JavaScript services make the appropriate API call to Hugging Face
5. JavaScript services forward the response back to the Go backend
6. Go backend forwards the response back to the client