# (DISCONTINUED) Lizbeth - Telegram App for Spotify Listening Session

Lizbeth is a Telegram bot developed to provide a collaborative music listening experience using Spotify. Lizbeth aimed to allow users to collaboratively listen to music and share their favorite tracks in real time via Telegram.

## Important Notice

**Update: Spotify has introduced an official feature for shared listening sessions.**

Spotify now provides a native feature for shared listening sessions directly within their app. This new feature offers enhanced functionality and seamless integration. For more details, please visit the [official Spotify support page](https://support.spotify.com/us/article/jam/).

## Features

- Collaborative listening to Spotify tracks in real-time
- Share and queue tracks within a Telegram group chat
- User-friendly commands for managing the session
- Integrated with Spotify's API for track information and playback

## Installation

### Prerequisites

- Go 1.16 or later
- A Spotify Developer Account
- A Telegram Bot Token

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/nenecchuu/lizbeth.git
   cd lizbeth
   ```

2. Set up environment variables:
   Create a `.env` file in the root directory with the following variables:
   ```
   SPOTIFY_CLIENT_ID=your_spotify_client_id
   SPOTIFY_CLIENT_SECRET=your_spotify_client_secret
   TELEGRAM_BOT_TOKEN=your_telegram_bot_token
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

4. Run the application:
   ```sh
   go run main.go
   ```

## Usage

1. Add the Telegram bot to your group chat.
2. Use the following commands to interact with the bot:
   - `/startsession` - Start a new listening session.
   - `/addtrack <track_name>` - Add a track to the session queue.
   - `/skip` - Skip the current track.
   - `/stopsession` - Stop the current session.

## Contributing

Since this project is discontinued, contributions are not actively sought. However, if you wish to fork and continue development, feel free to do so.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Go](https://golang.org/)
- [Spotify API](https://developer.spotify.com/documentation/web-api/)
- [Telegram Bot API](https://core.telegram.org/bots/api)

## Disclaimer

This project is not affiliated with or endorsed by Spotify. It was created as an independent project to provide a collaborative music listening experience.
