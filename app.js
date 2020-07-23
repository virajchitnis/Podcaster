const express = require('express');
const morgan = require('morgan');
const multer  = require('multer');
const podcast = require('podcast');
const xml = require('xml');
const fs = require('fs');
const path = require('path');

const config_dir = 'config';
const server_conf = require('./' + path.join(config_dir, 'server.json'));
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

// Build list of podcasts from configuration
const podcasts_dirs = fs.readdirSync(config_dir)
.filter(name => fs.lstatSync(path.join(config_dir, name)).isDirectory())
.map(name => './' + path.join(config_dir, name));

var podcasts = [];
podcasts_dirs.forEach((dir) => {
    var podcast_episodes = fs.readdirSync(dir)
    .filter(name => path.extname(name) === '.json' && name !== "info.json")
    .map(name => require('./' + path.join(dir, name)));

    podcasts.push({
        directory: dir,
        info: require('./' + path.join(dir, 'info.json')),
        episodes: podcast_episodes
    });
});

// Build endpoints for each podcast
podcasts.forEach((show) => {
    var feed = new podcast(show.info);
    show.episodes.forEach((episode) => {
        episode.enclosure.url = `${server_conf.websiteRoot}${episode.enclosure.url}`;
        feed.addItem(episode);
    });
    show.feedXML = feed.buildXml();

    app.get(`/${show.info.shortName}/feed.xml`, (req, res) => {
        res.set('Content-Type', 'text/xml');
        res.send(show.feedXML);
    });
});

app.listen(server_conf.port, () => console.log(`Podcaster listening at http://localhost:${server_conf.port}`));