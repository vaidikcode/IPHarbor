// huggingface-image.js - Service for handling Hugging Face image generation
import fetch from 'node-fetch';

/**
 * Handle image generation requests to Hugging Face
 * Supports two different API formats:
 * 1. Standard HF Inference API (model-specific endpoint)
 * 2. Nebius v1 API for image generations
 */
export async function handleHuggingFaceImage(req, res) {
  try {
    const { prompt, provider, model, content_type, url, token } = req.body;

    // Validate required inputs
    if (!prompt) {
      return res.status(400).json({ 
        error: 'Missing required parameter: prompt',
        success: false
      });
    }

    if (!token) {
      return res.status(400).json({ 
        error: 'Missing required parameter: token (API key)',
        success: false
      });
    }

    // Determine which API to use based on model or URL
    const isNebiusApi = url?.includes('nebius') || model?.includes('nebius');
    
    // Format the request data based on API type
    let data;
    let apiUrl;
    
    if (isNebiusApi) {
      // Format for Nebius API
      data = {
        response_format: "b64_json",
        prompt: prompt,
        model: model || "black-forest-labs/flux-dev"
      };
      apiUrl = url || "https://router.huggingface.co/nebius/v1/images/generations";
    } else {
      // Format for standard Hugging Face Inference API
      data = {
        inputs: prompt
      };
      
      // Determine URL: Use provided URL or construct from model
      if (url) {
        apiUrl = url;
      } else if (model) {
        apiUrl = `https://router.huggingface.co/hf-inference/models/${model}`;
      } else {
        apiUrl = "https://router.huggingface.co/hf-inference/models/black-forest-labs/FLUX.1-dev";
      }
    }

    console.log(`Making request to: ${apiUrl}`);
    console.log('With data:', JSON.stringify(data, null, 2));

    // Make request to the Hugging Face API
    const response = await fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });

    // Check for successful response
    if (!response.ok) {
      const errorText = await response.text();
      console.error('Error from Hugging Face API:', errorText);
      return res.status(response.status).json({
        error: `Error from Hugging Face API: ${response.statusText}`,
        details: errorText,
        success: false
      });
    }

    // Get response content type
    const contentType = response.headers.get('content-type');
    
    // Handle response based on content type
    if (contentType && contentType.includes('image')) {
      // For image responses, return the binary data
      const imageBuffer = await response.buffer();
      res.set('Content-Type', contentType);
      return res.send(imageBuffer);
    } else if (contentType && contentType.includes('application/json')) {
      // For JSON responses, parse and forward
      const jsonResponse = await response.json();
      return res.json(jsonResponse);
    } else {
      // For other responses, return as blob/binary
      const blobData = await response.buffer();
      res.set('Content-Type', contentType || 'application/octet-stream');
      return res.send(blobData);
    }
  } catch (error) {
    console.error('Error in handleHuggingFaceImage:', error);
    res.status(500).json({
      error: 'Failed to process image generation request',
      message: error.message,
      success: false
    });
  }
}
