# Youtube Playlist App

The application saves the metadata of __all__ the videos in __all__ the
playlists created by the authenticated user in a local text file.
For each video the data stored are the title, id, views, number of likes
and number of dislikes and name of file is suffixed by the current date.


## Use cases
- Videos are on Youtube are often removed without notice and their metadata are
lost. It is easy to recover them with the video id.

- If playlists are polled regularly and enough data are gathered, one can
generate more advanced statistics on them

- The current application can be modified/extended to store more metadata
possibly in a better storage format that allows queries and faster access.

## Setup

1. Add the [youtube API implementation library](https://github.com/googleapis/google-api-go-client)
to your Go installation with the following command

```
$ go get google.golang.org/api/youtube/v3
```

More info at https://github.com/googleapis/google-api-go-client/blob/master/GettingStarted.md


2. Setup the authorization

Follow the instructions in https://github.com/youtube/api-samples/blob/master/go/README.md
File oauth2.go was retrieved and should be configured based on the instructions there.

To run the code you need both a developer key (validates you as a developer) and
user authentication (anyone including yourself should give permission to the
application/script to access their data)