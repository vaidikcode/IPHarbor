// huggingface-video.js - Service for handling Hugging Face video generation
import fetch from 'node-fetch';
import { InferenceClient } from "@huggingface/inference";

/**
 * Handle video generation requests to Hugging Face
 * Supports different API formats:
 * 1. Standard HF Inference API (via fetch)
 * 2. InferenceClient API (via official HF client)
 * 3. Provider-specific APIs (novita, fal-ai, etc.)
 */
export async function handleHuggingFaceVideo(req, res) {
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

    // Use the InferenceClient for specific providers that require it
    if (provider === 'novita' || provider === 'fal-ai') {
      try {
        console.log(`Using HF InferenceClient for provider: ${provider}`);
        console.log(`Model: ${model || 'default'}, Prompt: ${prompt}`);
        
        const client = new InferenceClient(token);
        
        const video = await client.textToVideo({
          provider: provider,
          model: model || (provider === 'novita' ? 'Wan-AI/Wan2.1-T2V-14B' : 'Lightricks/LTX-Video'),
          inputs: prompt,
        });
        
        // Convert the Blob to a Buffer for sending
        const arrayBuffer = await video.arrayBuffer();
        const buffer = Buffer.from(arrayBuffer);
        
        res.set('Content-Type', 'video/mp4');
        return res.send(buffer);
      } catch (clientError) {
        console.error('Error using InferenceClient:', clientError);
        return res.status(500).json({
          error: 'Failed to generate video using InferenceClient',
          message: clientError.message,
          success: false
        });
      }
    }
    
    // Format the request data
    let data = {
      inputs: prompt
    };
    
    // Determine URL: Use provided URL or construct from model
    let apiUrl;
    if (url) {
      apiUrl = url;
    } else if (model) {
      apiUrl = `https://router.huggingface.co/hf-inference/models/${model}`;
    } else {
      apiUrl = "https://router.huggingface.co/hf-inference/models/Lightricks/LTX-Video";
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
    if (contentType && (contentType.includes('video') || contentType.includes('mp4'))) {
      // For video responses, return the binary data
      const videoBuffer = await response.buffer();
      res.set('Content-Type', contentType);
      return res.send(videoBuffer);
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
    console.error('Error in handleHuggingFaceVideo:', error);
    res.status(500).json({
      error: 'Failed to process video generation request',
      message: error.message,
      success: false
    });
  }
}
