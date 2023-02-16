# Improve roadmap

- [x] Remove uber.fx usage.
- [x] Refactor commands with cobra pkg.
- [x] Root config that aggregates all local configs.
- [x] Root repository
- [x] No application interfaces
- [ ] TestSuite Integration testing. 
- [x] Use squirrel as sql query builder instead of ent.
- [x] Context logger, as a wrapper for zap (just for joining values from upper scope to down).
- [ ] tracing spans for important places in code (almost everywhere).
- [ ] More interesting example (maybe url shortener)
- Prepare infrastructure packages.  
  - [ ] Web package with http.Client (prod. ready) 
  - [x] WebServer interface with Gin impl.
  - [ ] commands package with some impls (BaseCommand, StorageCommand)
  - [ ] storage package with Connectable interface (Connect, Disconnect)
  - [x] schema/json/ 
  - [x] Context logger with zap as engine
  - Any useful stuff that can be shared across others projects