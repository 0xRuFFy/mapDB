

GO = go

MAPDB = ./cmd/mapDB

mapDBLocal:
	$(GO) run $(MAPDB) -host=localhost