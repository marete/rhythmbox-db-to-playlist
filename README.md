# Introduction

`rhythmbox-db-to-playlist` is a simple utility to extract a playlist
usable by other players from the **Rhythmbox** song database, which is
normally stored as `rhythmdb.xml` in `~/.local/share/rhythmbox`. If
like me you have a large personal collection of music that you have
been playing with `rhythmbox` for a long time, then the database
contains useful historical data such as the play-count of the various
songs, time last played, first time played e.t.c.

You might want to extract a playlist in (say) reverse play-count order
and feed it to another media player. I use this utility to extract a
portable playlist in reverse play-count order, which playlist I can
play with some other player (e.g. `mpv`) when I am accessing my music
remotely for example by NFS share (since unfortunately Rhythmbox`s
buffering for slow network shares leads to many intermittent pauses
and is not usable over NFS or similar network shares (see [this
issue](https://gitlab.freedesktop.org/gstreamer/gst-plugins-base/-/issues/69)).

The tool can/will trivially be extended to export portable playlist in
other potentially interesting orders, for example by last-played or
first-seen order, if you would like to enjoy some music you last
played a decade ago ;)

# Usage

    cat ~/.local/share/rhythmbox/rhythmdb.xml | rhythmbox-db-to-playlist > playlist.txt
    mpv --playlist=playlist.txt

# Installation

    go install -v github.com/marete/rhythmbox-db-to-playlist
