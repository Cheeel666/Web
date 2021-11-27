This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        nginx/1.19.10
Server Hostname:        127.0.0.1
Server Port:            8082

Document Path:          /api/v1/courorts
Document Length:        52 bytes

Concurrency Level:      300
Time taken for tests:   179.318 seconds
Complete requests:      100000
Failed requests:        19354
   (Connect: 0, Receive: 0, Length: 19354, Exceptions: 0)
Non-2xx responses:      100000
Keep-Alive requests:    99989
Total transferred:      28293459 bytes
HTML transferred:       7251524 bytes
Requests per second:    557.67 [#/sec] (mean)
Time per request:       537.953 [ms] (mean)
Time per request:       1.793 [ms] (mean, across all concurrent requests)
Transfer rate:          154.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0      15
Processing:     0  492 2939.5      5   36381
Waiting:        0  492 2939.5      5   36381
Total:          0  492 2939.5      5   36381

Percentage of the requests served within a certain time (ms)
  50%      5
  66%     14
  75%     25
  80%     27
  90%     32
  95%     48
  98%   7329
  99%  22529
 100%  36381 (longest request)
