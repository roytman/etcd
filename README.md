# Benchmark etcd

A etcd benchmark tool based on fperf

# Installing

go get github.com/fperf/etcd/bin/fperf

# Usage

```
./fperf -server http://127.0.0.1:2379 -connection 256 -tick 1s etcd [put|get|delete|range]
```

The `key` is randomly generated with the default key-size = 4 bytes, so the key space is 2^32.