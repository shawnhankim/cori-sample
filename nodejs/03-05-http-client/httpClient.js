var http = require('http');

// Set information to call page
var options = {
    host: "127.0.0.1",
    port: 8000,
    path: "/",
};

// Call page
var req = http.request(options, (res) => {
    var data = "";

    // Receive data which is sent by server
    res.on('data', (chunk) => {
        data += chunk;
    });

    // Display data once it is completed to be received
    res.on('end', () => {
        console.log(data);
    });
});

// Complete request
req.end();
    
