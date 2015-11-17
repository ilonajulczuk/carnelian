# Carnelian

Carnelian, a simple irc bot.

Bot configuration and setup.
----------------------------

What should be configurable:
- nick
- list of channels
- commands that bot reacts to

Connecting to IRC
------------------

Bot connects to irc server.
Performs authorization.
Joins channels.

Responding to events.
---------------------

- ping events
- commands

Running.
--------

To run the bot:

    $ cd carnelian
    $ go build
    $ ./carnelian
