# trail-api

The trails behind the Trailhead app, served as JSON.

    GET /trails                    every trail
    GET /trails?difficulty=hard    trails of one difficulty
    GET /health                    ok when the service is up

## Deploying

Push to the `deploy` remote. The server builds and tests what it receives, and a push that
does not build is rejected with the output of the compiler.

    git push deploy main
