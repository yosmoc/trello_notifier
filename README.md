# Trello notifier

Notify current tasks that belongs in the specific list.

## dependency

### Linux

notify-send

### OSX

terminal-notifier

## Trello API key, token, list id

This program read following environmental variables

- TRELLO_API_KEY
- TRELLO_TOKEN
- TRELLO_LIST_ID

And if you put `.env` file in the same folder, it can automatically read this file.
