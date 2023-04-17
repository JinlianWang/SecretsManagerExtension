const http = require('http');

exports.handler = async (event) => {
  return new Promise((resolve, reject) => {
    const options = {
      hostname: 'localhost',
      port: 8080,
      path: '/ext-server',
      method: 'GET',
    };

    const req = http.request(options, (res) => {
      let responseBody = '';
      res.setEncoding('utf8');
      res.on('data', (chunk) => {
        responseBody += chunk;
      });
      res.on('end', () => {
        resolve({
          statusCode: 200,
          body: responseBody,
        });
      });
    });

    req.on('error', (error) => {
      reject(error);
    });

    req.end();
  });
};
