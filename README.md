# smoQerTT - A small fully functional MQTT broker

This project is basically a wrapper around the (Mochi MQTT)[https://github.com/mochi-mqtt/server] library.

It has the following features:

- Prometheus metrics listener to get server stats (right now not implemented its a wish :)
- Websocket listener using wss protocol to allow web browsers to connect
- TLS listener for backend clients

The reason why we are using this is because we wanted a very small mqtt broker that we can run in parallel with out api server. Its designed for small systems which do not need any scaling capabilities but rather just a simple, secure and
reliable broker.

[License: MIT](./LICENSE)
