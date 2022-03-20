# Mattermost Memes Plugin

[![Build Status](https://img.shields.io/circleci/project/github/mattermost/mattermost-plugin-memes/master.svg)](https://circleci.com/gh/mattermost/mattermost-plugin-memes)
[![Code Coverage](https://img.shields.io/codecov/c/github/mattermost/mattermost-plugin-memes/master.svg)](https://codecov.io/gh/mattermost/mattermost-plugin-memes)

**Maintainer:** [@hanzei](https://github.com/hanzei)

This plugin will create a slash command that you can use to create memes!

<img src="screenshot.png" width="583" height="426" />

`/meme everywhere "memes." "memes everywhere"`

For more information like available memes or command syntax type `/meme ` and press enter.

## Installation

From Mattermost 5.16 and later, the Memes Plugin is included in the Plugin Marketplace which can be accessed from **Main Menu > Plugins Marketplace**. You can install the Memes plugin there.

In Mattermost 5.15 and earlier, follow these steps:

1. Go to https://github.com/mattermost/mattermost-plugin-memes/releases/latest to download the latest release file in zip or tar.gz format.
2. Upload the file through **System Console > Plugins > Management**. See [documentation](https://docs.mattermost.com/administration/plugins.html#set-up-guide) for more details.

## Development

Read our documentation about the [Developer Workflow](https://developers.mattermost.com/extend/plugins/developer-workflow/) and [Developer Setup](https://developers.mattermost.com/extend/plugins/developer-setup/) for more information about developing and extending plugins.

Run `make memelibrary` to bundle up the meme assets (images, fonts, etc.).

For convenience, you can run the plugin from your terminal to preview an image for a given input. For example, on macOS, you can run the following to generate the above meme and open it in Preview:

`go run server/plugin.go -out demo.jpg 'memes. memes everywhere' && open demo.jpg`

This is especially useful when adding or modifying memes as you can quickly modify assets, `make memelibrary`, and preview the result using the above command. (See the files in ` memelibrary/assets` to get started with that.)

## Example commands
`go run server/plugin.go -out demo.jpg 'delete all the things!'`

Full arguments: Meme and double-quoted string arguments
`go run server/plugin.go -out demo.jpg 'sure-grandma "I remember when 56KBPS was the best" "sure grandma lets get you to bed"'`

Using a pattern: No string separated arguments, just pure pattern matching
`go run server/plugin.go -out demo.jpg 'I remember when 56kbps was the best sure grandma'`

`go run server/plugin.go -out demo.jpg 'boardroom "How do we improve EO?" "Develop faster" "Improve CI/CD" "Make memes on mattermost"'`
## Future Dev
Debug mode to draw the text boxes, and log out the current bounds
Add above as a command maybe?
Command line options to override font-colors, all_uppercase
Upload images for corporate meme?
Dynamic meme upload?