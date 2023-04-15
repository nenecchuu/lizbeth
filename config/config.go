package config

type MainConfig struct {
	Rest           RestConfig           `fig:"rest"`
	Mongo          MongoConfig          `fig:"mongo"`
	Spotify        SpotifyConfig        `fig:"spotify"`
	TelegramBotApi TelegramBotApiConfig `fig:"telegramBotApi"`
}

type (
	RestConfig struct {
		ListenAddress   string `fig:"listenAddress"`
		Port            string `fig:"port"`
		GracefulTimeout int    `fig:"gracefulTimeout"`
		AppName         string `fig:"appName"`
		ReadTimeout     int    `fig:"readTimeout"`
		WriteTimeout    int    `fig:"writeTimeout"`
		EnableSwagger   bool   `fig:"enableSwagger"`
	}
	MongoConfig struct {
		URI               string `yaml:"uri"`
		DB                string `yaml:"db"`
		ConnectionTimeout int    `yaml:"connectionTimeout"`
		PingTimeout       int    `yaml:"pingTimeout"`
	}
	// SECTION SPOTIFY
	SpotifyApiCredentials struct {
		AuthorizeCallbackUrl string `fig:"authorizeCallbackUrl"`
		Scope                string `fig:"scope"`
		ClientId             string `fig:"clientId"`
		ClientSecret         string `fig:"clientSecret"`
	}
	SpotifyHttpClientConfig struct {
		BaseUrl    string `fig:"baseUrl"`
		Timeout    int    `fig:"timeout"`
		RetryCount int    `fig:"retryCount"`
	}
	SpotifyCoreApiEndpoints struct {
		GetUserInfo string `fig:"get_user_info"`
		Search      string `fig:"search"`
		QueueTrack  string `fig:"queue_track"`
	}
	SpotifyAuthApiEndpoints struct {
		Authorize string `fig:"authorize"`
		Token     string `fig:"token"`
	}
	SpotifyCoreApiConfig struct {
		HttpClient SpotifyHttpClientConfig `fig:"httpConfig"`
		Endpoints  SpotifyCoreApiEndpoints `fig:"endpoints"`
	}
	SpotifyAuthApiConfig struct {
		HttpClient SpotifyHttpClientConfig `fig:"httpConfig"`
		Endpoints  SpotifyAuthApiEndpoints `fig:"endpoints"`
	}
	SpotifyConfig struct {
		Credentials SpotifyApiCredentials `fig:"credentials"`
		CoreApi     SpotifyCoreApiConfig  `fig:"coreApi"`
		AuthApi     SpotifyAuthApiConfig  `fig:"authApi"`
	}
	// END OF SPOTIFY

	// SECTION TELEGRAM BOT API
	TelegramBotApiConfig struct {
		Token        string `fig:"token"`
		DebugMode    bool   `fig:"debug"`
		Timeout      int    `fig:"timeout"`
		UpdateOffset int    `fig:"updateOffset"`
	}
	// END OF SECTION TELEGRAM BOT API
)
