// Import the WebSocket library
const WebSocket = require('ws');
const readline = require('readline');

// URL of your WebSocket server (replace with your actual server URL)
const wsUrl = 'ws://localhost:8080/api/ws';

// Create a new WebSocket connection
const ws = new WebSocket(wsUrl, {
    headers: {
        'Cookie': `jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk1MjkyODksImlhdCI6IjIwMjQtMTAtMThUMTg6NDg6MDkuMzcyNjYyMjIyKzAyOjAwIiwiaXNzIjoibm90ZXMtYXBwIiwic3ViIjoielF0aXpDNGhxaHdqM093SkMtMzkifQ.VB9bs4Ti7xXiZGTS3b6nwZFeWV8fA37YBCjZXRNl0Yk; Expires=Mon, 21 Oct 2024 16:48:09 GMT; HttpOnly; Secure; SameSite=Strict`
    }
});

// WebSocket connection opened
ws.on('open', function open() {
    console.log('Connected to WebSocket server');
    var i = 0;
    // Send a test message to the server
    while (i < 5) {
        i++;
        const message = JSON.stringify({
            type: 'message',
            content: 'my msg',
            timestamp: new Date().toISOString()
        });

        ws.send(message);
        console.log('Message sent:', message);
    }
});

let decoder = new TextDecoder("utf-8");

// Listen for messages from the server
ws.on('message', function incoming(data) {
    console.log('Message received from server:', decoder.decode(data));
});

// Handle WebSocket errors
ws.on('error', function(error) {
    console.error('WebSocket Error:', error);
});

// Handle WebSocket connection closure
ws.on('close', function(code, reason) {
    console.log(`WebSocket closed, code=${code}, reason=${reason}`);
});

