// server.js - Main entry point for the JavaScript services
import express from 'express';
import cors from 'cors';
import { handleHuggingFaceImage } from './services/huggingface-image.js';

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(cors());
app.use(express.json());

// Log all incoming requests
app.use((req, res, next) => {
  console.log(`${new Date().toISOString()} - ${req.method} ${req.path}`);
  next();
});

// Routes
app.post('/huggingface/image', handleHuggingFaceImage);

// Error handling middleware
app.use((err, req, res, next) => {
  console.error('Error:', err);
  res.status(500).json({ 
    error: 'Internal Server Error', 
    message: err.message,
    success: false 
  });
});

// Start server
app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
  console.log('Available endpoints:');
  console.log('- POST /huggingface/image - For image generation');
});
