This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        nginx/1.19.10
Server Hostname:        127.0.0.1
Server Port:            8082

Document Path:          /api/v1/courorts
Document Length:        52 bytes

Concurrency Level:      10
Time taken for tests:   0.855 seconds
Complete requests:      5000
Failed requests:        0
Non-2xx responses:      5000
Total transferred:      1350000 bytes
HTML transferred:       260000 bytes
Requests per second:    5851.33 [#/sec] (mean)
Time per request:       1.709 [ms] (mean)
Time per request:       0.171 [ms] (mean, across all concurrent requests)
Transfer rate:          1542.83 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0      29
Processing:     1    1   1.4      1      30
Waiting:        1    1   1.4      1      30
Total:          1    2   1.5      1      30

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      2
  75%      2
  80%      2
  90%      2
  95%      3
  98%      4
  99%      6
 100%     30 (longest request)
