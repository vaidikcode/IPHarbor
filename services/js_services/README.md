# IPHarbor JavaScript Services

This folder contains JavaScript services that handle requests from the Go backend and forward them to various AI providers.

## Services

### Hugging Face Image Generation

Handles image generation requests to Hugging Face. Supports two API formats:

1. **Standard Hugging Face Inference API**
   ```javascript
   // Example API call
   const response = await fetch(
     "https://router.huggingface.co/hf-inference/models/black-forest-labs/FLUX.1-dev",
     {
       headers: {
         Authorization: "Bearer hf_xxxxxxxxxxxxxxxxxxxxxxxx",
         "Content-Type": "application/json",
       },
       method: "POST",
       body: JSON.stringify({ inputs: "Astronaut riding a horse" }),
     }
   );
   ```

2. **Nebius v1 API for Image Generation**
   ```javascript
   // Example API call
   const response = await fetch(
     "https://router.huggingface.co/nebius/v1/images/generations",
     {
       headers: {
         Authorization: "Bearer hf_xxxxxxxxxxxxxxxxxxxxxxxx",
         "Content-Type": "application/json",
       },
       method: "POST",
       body: JSON.stringify({
         response_format: "b64_json",
         prompt: "Astronaut riding a horse",
         model: "black-forest-labs/flux-dev",
       }),
     }
   );
   ```

## Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the server:
   ```bash
   npm start
   ```

The server will run on port 3000 by default.

## API Endpoints

### POST /huggingface/image

Generates an image using Hugging Face models.

**Request Format:**
```json
{
  "prompt": "Your text prompt here",
  "provider": "huggingface",
  "model": "black-forest-labs/FLUX.1-dev",
  "content_type": "application/json",
  "url": "optional-direct-url",
  "token": "your-api-token"
}
```

**Response:**
- For successful requests: Image binary data with appropriate Content-Type
- For errors: JSON error response

## How It Works

1. The Go backend receives a request and forwards it to this service
2. This service analyzes the request parameters to determine which API to use
3. It constructs the appropriate API call to Hugging Face
4. It forwards the response back to the Go backend
