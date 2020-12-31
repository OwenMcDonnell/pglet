const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = app => {
  app.use("/ws", createProxyMiddleware({target: "http://localhost:5000", ws: true}))
}