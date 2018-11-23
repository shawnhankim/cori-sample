// Load net module
var net = require('net');

// Create TCP server and reply "hello world" once client is connected
var svr = net.createServer((socket) => {
    socket.end("hello world");
});

// Handle network error
svr.on('error', (err) => {
    console.log(err);
});

// Listen 9000 port
svr.listen(9000, () => {
    // Display message once listening is be able to be accepted
    console.log('listen', svr.address());
});


