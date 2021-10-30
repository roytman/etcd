# Etcd long transaction tests

The test environment:


## Results
log level - info, requests were not sent to the leader 
Transaction length  Min time    Average time    Max time 
500                 5.75 ms     9.08 ms         25.07 ms
1000                11.18 ms    14.7 ms         42.79 ms
2000                23.49 ms    27.57 ms        63.05 ms
4000                46.51 ms    53.23 ms        100.65 ms
8000                95.02 ms    215.87 ms       1.564 s
16000               1.825 s     3.273 s         4.662 s
32000               4.957 s     4.957 s         

log level - warn requests were sent to the leader
Transaction length  Min time    Average time    Max time 
500                 
1000                11.32 ms    14.32  ms       46.88  ms
2000                23.78 ms    27.595 ms       80.371 ms
4000                47.24 ms    53.703 ms       187.488 ms
8000                93.465 ms   117.78 ms       787.735 ms
16000               319.32 ms   804.3 ms        1.347 s
32000               1.482 s     1.818 s         2.084 s

--experimental-txn-mode-write-with-shared-buffer=false sent to leader
Transaction length  Min time    Average time    Max time 
500                 5.230 ms    8.00 ms         27.00 ms
1000                11.07 ms    13.82 ms        34.39 ms
2000                23.03 ms    26.79 ms        61.11 ms
4000                45.80 ms    52.57 ms        95.39 ms
8000                91.9 ms     109.01 ms       224.43 ms
16000               304.01 ms   1.2 s           3.98 s
32000               1.46 s      1.95 s          5.05 s

--experimental-txn-mode-write-with-shared-buffer=true sent to leader
Transaction length  Min time    Average time    Max time 
500                 5.17 ms     7.66 ms         25.11 ms
1000                11.01 ms    13.46 ms        35.21 ms
2000                22.51 ms    26.34 ms        69.78 ms
4000                46.05 ms    51.97 ms        88.63 ms
8000                90.76 ms    111.09 ms       210.0 ms
16000               661.99 ms   1.05 s          2.20 s
32000               1.52 s      2.06 s          3.95 s

--experimental-txn-mode-write-with-shared-buffer=true sent to leader, log-level=error
Transaction length  Min time    Average time    Max time
500                 5.18 ms     7.61 ms         37.23 ms
1000                10.59 ms    13.42 ms        29.73 ms         
2000                21.77 ms    26.10 ms        52.54 ms    
4000                45.52 ms    51.68 ms        80.52 ms
8000                91.43 ms    107.27 ms       222.72 ms
16000               281.65 ms   386.47 ms       494.54 ms
32000               720.71 ms   794.12 ms       968.59 ms
64000               1.596 s     1.742 s         1.902 s
128000              3.295 s     3.493 s         3.78 s
256000                          7.5 s