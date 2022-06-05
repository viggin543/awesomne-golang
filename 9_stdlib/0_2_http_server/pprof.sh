# use the pprof tool to look at the heap profile:
go tool pprof http://localhost:8080/debug/pprof/heap
# then type "web"
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
# then type "web"