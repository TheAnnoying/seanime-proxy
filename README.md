This is a fork of [Seanime](https://github.com/5rahim/seanime) that adds support for using a custom server URL instead of the local server hosted by Denshi.

This feature is described in [this issue](https://github.com/5rahim/seanime/issues/854).

## Usage

After installation, launch Denshi once so it generates a `config.toml` file.

If you have previously installed Seanime, you should already have a `config.toml`, so this step is unnecessary.

### `config.toml` location

* **Windows:** `%APPDATA%\Seanime\config.toml`
* **Linux:** `$XDG_CONFIG_HOME/Seanime/config.toml` or `$HOME/.config/Seanime/config.toml`
* **macOS:** `$HOME/Library/Application Support/Seanime/config.toml`

### Configure a custom server

Open `config.toml` and find the existing `[server]` section. Do not remove any of the other settings in this section. Simply add or modify the `proxy_server_url` entry:

```toml
[server]
# other existing settings...
proxy_server_url = 'http://myserver:3211'
# other existing settings...
```

Replace `http://myserver:3211` with the URL of the server you want Denshi to use.

Once you launch Denshi, it should connect to your specified server.
