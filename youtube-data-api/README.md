# YouTube Data API
## Requirements
- Golang
    - github.com/joho/godotenv
    - google.golang.org/api/googleapi/transport
    - google.golang.org/api/youtube/v3
- YouTube Data API
- Google OAuth Client ID
    - To use insert or reply comment

## Usage in Golang
1. Clone code
```
$ git clone https://github.com/Doarakko/api-challenge
$ cd youtube-data-api
```

2. Enter your key, id, etc
```
$ mv .env.example .env
```
```
YOUTUBE_API_KEY = abcdef
CLIENT_ID = ghijk.apps.googleusercontent.com
CLIENT_SECRET = lmnopqr
```

3. Edit main function
- Print channel detail information

```
func main() {
    err := godotenv.Load("./.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    printChannelInfo("channel id")
}
```

- Print video comments
```
printComments(getComments("video id"))
```