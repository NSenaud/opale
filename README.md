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


CLI tool
--------

```
opale-cli get cpu --percentage
```


Required Elements to Operate
----------------------------

- [x] Data loading
- [x] SQLite saving
- [x] gRPC communication
- [x] TOML configuration file
- [x] Use XDG path
- [x] Leveled logging
- [x] CLI parameters
- [ ] SQLite cleanup entries
- [x] SQLite request


Available Monitors
------------------

- [x] CPU
- [x] RAM
- [ ] Swap
- [ ] Load Average
- [ ] Network i/o
- [ ] Storage i/o
- [ ] Storage status
- [ ] Battery
- [ ] Screen backlight
- [ ] Temperatures
- [ ] Network settings (IP, hostname, SSID...)
- [ ] Process count
- [ ] Kernel version


Features
--------

- [x] Server:
  - [x] Monitor items
  - [x] Save them in SQLite database
  - [x] TOML configuration file
  - [x] Customizable items
  - [ ] Customizable retention policy
  - [ ] Customizable saving interval
- [x] Simple query client:
  - [x] Query server for a given item
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
