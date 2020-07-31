Podcaster
=========

Host your own podcasts on your server.

Intro
-----

My brother recently started podcasting, and in order to help publish his podcasts, I created a Go (golang) application that runs on my server.
I have decided to create an open source version of this application and release it here. 

Endpoints
---------

- `/health` - Simply returns a `200 OK` response with a `Up!` message as the body. This can be used by frontend web servers such as Haproxy to check whether this server is running.
- `/{podcast_name}/feed.xml` - This is the feed for a configured podcast. This is what will be used by podcasting applications, or by Apple, Spotify, etc.

Setup
-----

1. Download the relevant binary file from the releases section of this repository.
2. Run the binary file from the command line interface of your server: `./Podcaster`. Run it with the `-config "path/to/config.yaml"` option if you want to specify a config file. The default location for the config file is: `/etc/podcaster/config.yaml`. If no config file is found at the default location (or at the location specified), one will be automatically created on first run.

Dockerize
---------

To run this application within a docker container:

1. Build the docker image: `docker build -t Podcaster .`
2. Run a docker container: `docker run --name Podcaster --restart always -v "path/to/config/directory":/etc/podcaster -v "path/to/data/directory":/var/podcaster -d Podcaster`
3. Restart the docker container if you make changes to the data directory later: `docker restart Podcaster`

Important
---------

- Make sure the `shortname` field in the `.yaml` file for each podcast does not contain spaces or special characters. This will be used in the URL for the podcast feed.
- `href` field must contain the location of your podcast's cover art. This location must be relative to the data directory. For instance if you want the URL for the cover art to be `https://www.domain.tld/media/images/podcast_name.jpeg`, this field must contain: `/media/images/podcast_name.jpeg` and the file should be stored in `data/directory/media/images/podcast_name.jpg`.
- `url` field in the episode configuration must meet the same requirements as the `href` field mentioned above.
- `file` field in the episode configuration must contain the on disk location of your podcast media (audio/video) files relative to the data directory. For example, if you want to serve the media file from `https://www.domain.tld/media/podcast_name/s01e01.mp3`, you should store it in the `data/directory -> media -> podcast_name` directory and name it `s01e01.mp3`. You must then put `/media/podcast_name/s01e01.mp3` as the value for the `file` field.

Contributing
------------

If you want to suggest a change or fix something, feel free to send me a pull request.

License
-------

This project is licensed under the GNU GPLv3 license.