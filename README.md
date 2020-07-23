Podcaster
=========

Host your own podcasts on your server.

Intro
-----

My brother recently started podcasting, and in order to help publish his podcasts, I created a nodejs application that runs on my server.
I have decided to create an open source version of this application and release it here. 

Setup
-----

1. Clone this repository.
2. Run the `setup.sh` script.
3. Setup the configuration directory. Refer to the `example_config` directory in the project. You can rename it to `config` and edit it. The `setup.sh` script will copy it for you.
4. Place your audio/video podcast files and cover art image files in the `public` directory. You can have sub-directories for each type of media or for each podcast.
5. Start the server: `npm start`.

Important
---------

- Make sure the `shortName` field in the `info.json` for each podcast does not contain spaces or special characters. This will be used in the URL for the podcast feed.
- `itunesImage` and `image_url` fields must contain the location of your podcast's cover art. This location must be relative to this project. For instance if you want the URL for the cover art to be `https://www.domain.tld/media/images/podcast_name.jpeg`, these two fields must contain: `/media/images/podcast_name.jpeg`.
- `url` field in the episode configuration must meet the same requirements as the `itunesImage` and `image_url` fields above.
- `file` field in the episode configuration must contain the on disk location of your podcast media (audio/video) files relative to the root of this project. For example, if you want to server the media file from `https://www.domain.tld/media/podcast_name/s01e01.mp3`, you should store it in the `public -> media -> podcast_name` directory in this project and and name it `s01e01.mp3`. You must then put `./public/media/podcast_name/s01e01.mp3` as the value for the `file` field.

Contributing
------------

If you want to suggest a change or fix something, feel free to send me a pull request.

License
-------

This project is licensed under the GNU GPLv3 license.