# Improve roadmap

- [x] Remove uber.fx usage.
- [x] Refactor commands with cobra pkg.
- [x] Root config that aggregates all local configs.
- [x] Root repository
- [x] No application interfaces
- [ ] TestSuite Integration testing with less mocks (because no Interfaces). 
- [x] Use squirrel as sql query builder instead of ent.
- [x] Context logger, as a wrapper for zap (just for joining values from upper scope to down).
- [ ] tracing spans for important locations in code (almost everywhere).
- [ ] More complicated, real world example
- Prepare infrastructure packages.  
  - [ ] Web package with http.Client (prod. ready) WebServer interface with Gin impl.
  - [ ] commands package with some impls (BaseCommand, StorageCommand)
  - [ ] storage package with Connectable interface (Connect, Disconnect)
  - [ ] schema/json/ 
  - [ ] Context logger with zap as engine
  - [ ] etc