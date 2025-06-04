# IPHarbor Services

This folder contains service implementations that handle requests from the Go backend and forward them to various AI providers.

---

## API Documentation

### Centralized Service Forwarder (Go Agent)

The Go backend acts as a centralized service forwarder that receives requests from clients and forwards them to the appropriate AI service endpoints. When registering AI agents, provide us with your hosting provider, model name, content type, service endpoint, and token to access this. The prompt will be from the user.

#### Input Format

All endpoints accept the following JSON structure:

```json
{
  "prompt": "Your text prompt here",
  "provider": "huggingface",
  "model": "model-name",
  "content_type": "application/json",
  "url": "optional-direct-url",
  "token": "your-api-token"
}
```

#### Endpoints

1. **Image Generation**
   - Endpoint: `/generateTextToImage`
   - Forwards to: `http://localhost:3000/huggingface/image`
   - Example model: "black-forest-labs/FLUX.1-dev"

2. **Text Generation**
   - Endpoint: `/generateTextToText`
   - Not yet implemented

3. **Audio Processing**
   - Endpoint: `/generateTextToAudio`
   - Not yet implemented

4. **Video Generation**
   - Endpoint: `/generateTextToVideo`
   - Not yet implemented

#### Output Format

The service forwards the exact response received from the destination service, including:
- Status code
- Headers
- Content body (which may be JSON, binary image/video data, or another format)

### JavaScript Services (js_services)

The JavaScript services handle requests from the Go backend and forward them to various AI providers. They analyze the request parameters to determine which API to use and construct the appropriate API call.

#### Available Services

1. **Hugging Face Image Generation** - `/huggingface/image`
   - Supports two API formats:
     - Standard Hugging Face Inference API
     - Nebius v1 API for Image Generation

For more details, see the [JavaScript Services README](./js_services/README.md).