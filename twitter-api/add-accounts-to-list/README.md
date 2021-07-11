# add accounts to list

Add accounts who tweet specific keyword to twitter list.

## Requirements

- go
- Twitter API

## Usage

1. Setup

```sh
git clone https://github.com/Doarakko/twitter-q-listing
cd twitter-q-listing
cp .env.example .env
```

2. Enter your environment variables in `.env`

- `LIST_MODE`

  - Scope of list.
  - Set `private` or `public`

- `QUERY`

  - Keyword to search

- `QUERY_HASHTAG`

  - Add `#` to `Query`
  - Set `true` or `false`

- `TWEET_TYPE`
  - `popular`
  - `recent`
  - `mixed`: Mix `popular` and `recent`

3. Run

```sh
go run main.go
```

## Author

Doarakko

## License

MIT
