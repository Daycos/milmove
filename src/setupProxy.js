const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = (app) => {
  app.use(createProxyMiddleware('/api', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/internal', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/admin', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/ghc', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/prime', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/support', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/testharness', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/storage', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/devlocal-auth', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/auth/**', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/logout', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/downloads', { target: 'http://milmove.daycos.com:8080/' }));
  app.use(createProxyMiddleware('/debug/**', { target: 'http://milmove.daycos.com:8080/' }));
};
