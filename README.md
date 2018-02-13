# GraphDb

There are a few packages in this project.

## GraphDb

The basic on is a simple Graph Database built on the
[stow](https://github.com/graymeta/stow) project.

The graph database provides a way to store and retrieve nodes with edges.

## DepDb

DepDb is a dependency database. This uses the graph database to
represent dependencies between events. An event is abstract, but can be
expressed as a match to a regular expression or a basic type of event.

This package is incomplete, but is a start.

There is a timeout concept such that an event could be auto triggered
after a period of inactivity. As with most good bits, this is not implemented.
