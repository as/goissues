# Summary

- The main benchmarking routine executes 2 `Get`s (one in the background).
- Each writer executes a `Put` with 1/8 probability (i.e., an occasion).
- For each writer, there are an additional `8` concurrent readers

```
; go test -bench .
goos: linux
goarch: amd64
pkg: github.com/as/goissues/28938
BenchmarkMap/0W2R/Sync/Incr/Mask0-4         	30000000	        41.1 ns/op
BenchmarkMap/0W2R/Sync/Incr/Maskf-4         	30000000	        42.0 ns/op
BenchmarkMap/0W2R/Sync/Incr/Maskff-4        	30000000	        48.6 ns/op
BenchmarkMap/0W2R/Sync/Incr/Maskfff-4       	20000000	        71.8 ns/op
BenchmarkMap/0W2R/Sync/Incr/Maskffff-4      	10000000	       204 ns/op
BenchmarkMap/0W2R/Sync/Incr/Maskfffff-4     	10000000	       235 ns/op
BenchmarkMap/0W2R/Sync/Rand/Mask0-4         	20000000	        57.6 ns/op
BenchmarkMap/0W2R/Sync/Rand/Maskf-4         	20000000	        57.7 ns/op
BenchmarkMap/0W2R/Sync/Rand/Maskff-4        	20000000	        67.1 ns/op
BenchmarkMap/0W2R/Sync/Rand/Maskfff-4       	20000000	       109 ns/op
BenchmarkMap/0W2R/Sync/Rand/Maskffff-4      	 3000000	       402 ns/op
BenchmarkMap/0W2R/Sync/Rand/Maskfffff-4     	 3000000	       537 ns/op
BenchmarkMap/0W2R/Lock/Incr/Mask0-4         	20000000	        98.7 ns/op
BenchmarkMap/0W2R/Lock/Incr/Maskf-4         	20000000	        99.4 ns/op
BenchmarkMap/0W2R/Lock/Incr/Maskff-4        	20000000	       103 ns/op
BenchmarkMap/0W2R/Lock/Incr/Maskfff-4       	10000000	       116 ns/op
BenchmarkMap/0W2R/Lock/Incr/Maskffff-4      	10000000	       185 ns/op
BenchmarkMap/0W2R/Lock/Incr/Maskfffff-4     	10000000	       202 ns/op
BenchmarkMap/0W2R/Lock/Rand/Mask0-4         	20000000	       109 ns/op
BenchmarkMap/0W2R/Lock/Rand/Maskf-4         	20000000	       110 ns/op
BenchmarkMap/0W2R/Lock/Rand/Maskff-4        	20000000	       109 ns/op
BenchmarkMap/0W2R/Lock/Rand/Maskfff-4       	20000000	       117 ns/op
BenchmarkMap/0W2R/Lock/Rand/Maskffff-4      	10000000	       203 ns/op
BenchmarkMap/0W2R/Lock/Rand/Maskfffff-4     	10000000	       233 ns/op
BenchmarkMap/1W10R/Lock/Incr/Mask0-4        	 1000000	      2046 ns/op
BenchmarkMap/1W10R/Lock/Incr/Maskf-4        	 1000000	      1999 ns/op
BenchmarkMap/1W10R/Lock/Incr/Maskff-4       	  500000	      2038 ns/op
BenchmarkMap/1W10R/Lock/Incr/Maskfff-4      	  500000	      2872 ns/op
BenchmarkMap/1W10R/Lock/Incr/Maskffff-4     	  500000	      2765 ns/op
BenchmarkMap/1W10R/Lock/Incr/Maskfffff-4    	 1000000	      2854 ns/op
BenchmarkMap/1W10R/Lock/Rand/Mask0-4        	 1000000	      2480 ns/op
BenchmarkMap/1W10R/Lock/Rand/Maskf-4        	 1000000	      2554 ns/op
BenchmarkMap/1W10R/Lock/Rand/Maskff-4       	 1000000	      2694 ns/op
BenchmarkMap/1W10R/Lock/Rand/Maskfff-4      	  500000	      3393 ns/op
BenchmarkMap/1W10R/Lock/Rand/Maskffff-4     	  500000	      3161 ns/op
BenchmarkMap/1W10R/Lock/Rand/Maskfffff-4    	  300000	      3682 ns/op
BenchmarkMap/1W10R/Sync/Incr/Mask0-4        	10000000	       242 ns/op
BenchmarkMap/1W10R/Sync/Incr/Maskf-4        	 5000000	       312 ns/op
BenchmarkMap/1W10R/Sync/Incr/Maskff-4       	 5000000	       343 ns/op
BenchmarkMap/1W10R/Sync/Incr/Maskfff-4      	 1000000	      1114 ns/op
BenchmarkMap/1W10R/Sync/Incr/Maskffff-4     	 2000000	      1210 ns/op
BenchmarkMap/1W10R/Sync/Incr/Maskfffff-4    	 1000000	      1307 ns/op
BenchmarkMap/1W10R/Sync/Rand/Mask0-4        	 5000000	       415 ns/op
BenchmarkMap/1W10R/Sync/Rand/Maskf-4        	 3000000	       385 ns/op
BenchmarkMap/1W10R/Sync/Rand/Maskff-4       	 3000000	       588 ns/op
BenchmarkMap/1W10R/Sync/Rand/Maskfff-4      	 1000000	      1056 ns/op
BenchmarkMap/1W10R/Sync/Rand/Maskffff-4     	 1000000	      1748 ns/op
BenchmarkMap/1W10R/Sync/Rand/Maskfffff-4    	 1000000	      2027 ns/op
PASS
ok  	github.com/as/goissues/28938	387.041s
```