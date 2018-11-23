var net = require('net');

// Set connection information
var options = {
    port: 9000,
    host: "127.0.0.1"
};

// Connect server
var client = net.connect(options, () => {
    console.log("connected to server");
});

// Event to receive data
client.on('data', (data) => {
    console.log(data.toString());
});

// Event to disconnect server
client.on('end', () => {
    console.log("disconnected");
});
