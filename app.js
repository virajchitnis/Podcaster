const express = require('express');
const morgan = require('morgan');
const multer  = require('multer');
const podcast = require('podcast');
const xml = require('xml');
const fs = require('fs');
const path = require('path');

const server_conf = require('./config/server.json');
const app = express();
app.use(morgan(server_conf.logStyle));
app.use(express.static('public'));

/* Endpoint for checking server health.
 * This can be used by a frontend server such as Haproxy to check if
 * this server is up or not. 
 */
app.get('/health', (req, res) => {
    res.send('Up!');
});

app.listen(server_conf.port, () => console.log(`Podcaster listening at http://localhost:${server_conf.port}`));