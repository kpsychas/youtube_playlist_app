# Youtube Playlist App

The application saves the metadata of __all__ the videos in one or more of the
playlists created by the authenticated user.
For each video the data stored are the title, id, views, number of likes
and number of dislikes and name of file is suffixed by the current date.


## Use cases
- Videos are on Youtube are often removed without notice and their metadata are
lost. With this application it is easy to find them from local storage by
searching their video id.

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

## Run the script

Include dependencies in run command as

```
go run .\youtube_playlists.go .\errors.go .\oauth2.go
```
