#!/usr/bin/env bash
docker build . -t dslic_image
heroku container:push web --app dslic
heroku container:release web --app dslic
echo deployment done.
