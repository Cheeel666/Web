This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        nginx/1.19.10
Server Hostname:        127.0.0.1
Server Port:            8082

Document Path:          /api/v1/courorts
Document Length:        52 bytes

Concurrency Level:      8
Time taken for tests:   1.034 seconds
Complete requests:      5000
Failed requests:        0
Non-2xx responses:      5000
Total transferred:      1270000 bytes
HTML transferred:       260000 bytes
Requests per second:    4835.53 [#/sec] (mean)
Time per request:       1.654 [ms] (mean)
Time per request:       0.207 [ms] (mean, across all concurrent requests)
Transfer rate:          1199.44 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       4
Processing:     0    1   1.0      1      12
Waiting:        0    1   1.0      1      12
Total:          0    2   1.1      1      12

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      2
  80%      2
  90%      3
  95%      4
  98%      6
  99%      7
 100%     12 (longest request)
