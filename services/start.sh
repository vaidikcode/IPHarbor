#!/bin/bash
# Start the JavaScript services

cd "$(dirname "$0")/js_services"

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "Node.js is not installed. Please install Node.js to run the JavaScript services."
    exit 1
fi

# Install dependencies if needed
if [ ! -d "node_modules" ]; then
    echo "Installing dependencies..."
    npm install
fi

# Start the server
echo "Starting JavaScript services on port 3000..."
npm start
