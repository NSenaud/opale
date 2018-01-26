Opale
=====


Simple Linux monitoring with history saving.


What Opale is
-------------

A Linux laptop and desktop monitoring utility with history saving in a local
SQLite database. The utility is build around a client-server architecture, with
the server saving monitored items (CPU load, memory usage, battery status...)
into the database at a given interval, and answering to clients request. A
client can be used to display items into a window manager status bar, or to
show a graph of values. Watchers could also be implemented to send alerts or
notifications or some pre-defined conditions.


What Opale is not
-----------------

Opale is not a server-monitoring solution, there are already plenty of these
which are battle-tested and available for free. Think more of Opale as the
Linux version of [iStat][istat].


Features
--------

- [ ] Server:
  - [ ] Monitor items
  - [ ] Save them in SQLite database
  - [ ] TOML configuration file
  - [ ] Customizable items
  - [ ] Customizable retention policy
  - [ ] Customizable saving interval
- [ ] Simple query client:
  - [ ] Query server for a given item
  - [ ] Customizable answer format
- [ ] Curses client:
  - [ ] Show graph of item values
  - [ ] Customizable date range
- [ ] Notification watcher:
  - [ ] Be alerted via native notifications of an event
  - [ ] Advanced customization (rules, histeresis...)
- [ ] PushOver watcher:
  - [ ] Be alerted via PushOver of events


[istat]: https://bjango.com/mac/istatmenus/
