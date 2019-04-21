package main

import (
        //"encoding/json"
        "fmt"
        "log"
        "os"
        "strconv"
        "time"

        "google.golang.org/api/youtube/v3"
)


// Retrieve playlistItems in the specified playlist
func playlistItemsList(service *youtube.Service, part string, playlistId string, pageToken string) *youtube.PlaylistItemListResponse {
        call := service.PlaylistItems.List(part)
        call = call.PlaylistId(playlistId)
        if pageToken != "" {
                call = call.PageToken(pageToken)
        }
        response, err := call.Do()
        handleError(err, "")
        return response
}

// Retrieve resource for the authenticated user's playlists
func playlistsListMine(service *youtube.Service, part string, pageToken string) *youtube.PlaylistListResponse {
        call := service.Playlists.List(part)
        call = call.Mine(true)
        if pageToken != "" {
                call = call.PageToken(pageToken)
        }
        response, err := call.Do()
        handleError(err, "")
        return response
}

// Get single video
func videoList(service *youtube.Service, part string, id string) *youtube.VideoListResponse {
        call := service.Videos.List(part)
        call = call.Id(id)
        response, err := call.Do()
        handleError(err, "")
        return response
}


func printPlaylistVideos(service *youtube.Service, playlistId string, f *os.File) {
    nextPageToken := ""
    for {
        // Retrieve next set of items in the playlist.
        playlistResponse := playlistItemsList(service, "snippet", playlistId, nextPageToken)

        for _, playlistItem := range playlistResponse.Items {
                title := playlistItem.Snippet.Title
                videoId := playlistItem.Snippet.ResourceId.VideoId
                videoResponse := videoList(service, "statistics", videoId)

                // If video is removed/private, statistics will be empty
                viewCount := ""
                likeCount := ""
                dislikeCount := ""
                for _, videoItem := range videoResponse.Items {
                    viewCount = strconv.Itoa(int(videoItem.Statistics.ViewCount))
                    likeCount = strconv.Itoa(int(videoItem.Statistics.LikeCount))
                    dislikeCount = strconv.Itoa(int(videoItem.Statistics.DislikeCount))
                }
                fmt.Fprintf(f, "%v, %v, %v, %v, %v \r\n", title, videoId,
                    viewCount, likeCount, dislikeCount)

        }

        // Set the token to retrieve the next page of results
        // or exit the loop if all results have been retrieved.
        nextPageToken = playlistResponse.NextPageToken
        if nextPageToken == "" {
                break
        }
    }
}


func main() {
    // Get Client
    client := getClient(youtube.YoutubeReadonlyScope)
    service, err := youtube.New(client)

    if err != nil {
            log.Fatalf("Error creating YouTube client: %v", err)
    }

    // Create and open file for results
    t := time.Now()
    outputFileName := t.Format("./results2006-01-02.txt")
    f, err := os.Create(outputFileName)

    if err != nil {
            log.Fatalf("Error creating file: %v", err)
    }

    defer f.Close()

    nextPageToken := ""

    for {
        response := playlistsListMine(service, "snippet", nextPageToken)


        for _, playlist := range response.Items {
            playlistId := playlist.Id
            playlistTitle := playlist.Snippet.Title

            fmt.Fprintf(f, "================================\r\n")
            fmt.Fprintf(f, "Videos in list %s\r\n", playlistTitle)

            printPlaylistVideos(service, playlistId, f)
        }

        nextPageToken = response.NextPageToken
        if nextPageToken == "" {
                break
        }
    }
}