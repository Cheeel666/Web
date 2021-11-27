ab -n 100000 -c 400 -k http://127.0.0.1:8082/api/v1/courorts/ > /Users/ilchel/go/src/git/Web/frontend/nginx/logs/res.md

# С балансировкой
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        SkOpen
Server Hostname:        127.0.0.1
Server Port:            8082

Document Path:          /api/v1/courorts/
Document Length:        122 bytes

Concurrency Level:      400
Time taken for tests:   40.201 seconds
Complete requests:      100000
Failed requests:        89425
   (Connect: 0, Receive: 0, Length: 89425, Exceptions: 0)
Non-2xx responses:      87098
Keep-Alive requests:    100000
Total transferred:      32646222 bytes
HTML transferred:       15060942 bytes
Requests per second:    2487.51 [#/sec] (mean)
Time per request:       160.803 [ms] (mean)
Time per request:       0.402 [ms] (mean, across all concurrent requests)
Transfer rate:          793.05 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.8      0      19
Processing:     0  158 1684.6     10   30531
Waiting:        0  158 1684.6     10   30531
Total:          0  158 1684.6     10   30531

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     11
  75%     11
  80%     11
  90%     16
  95%    106
  98%    768
  99%   1633
 100%  30531 (longest request)
```


# без балансировки 

```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)


Server Software:        SkOpen
Server Hostname:        127.0.0.1
Server Port:            8082

Document Path:          /api/v1/courorts/
Document Length:        122 bytes

Concurrency Level:      400
Time taken for tests:   323.641 seconds
Complete requests:      100000
Failed requests:        24467
   (Connect: 0, Receive: 0, Length: 24467, Exceptions: 0)
Non-2xx responses:      6937
Keep-Alive requests:    100000
Total transferred:      53347103 bytes
HTML transferred:       10381192 bytes
Requests per second:    308.98 [#/sec] (mean)
Time per request:       1294.563 [ms] (mean)
Time per request:       3.236 [ms] (mean, across all concurrent requests)
Transfer rate:          160.97 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0      15
Processing:     0 1274 3673.0     65   30121
Waiting:        0 1274 3673.0     65   30121
Total:          0 1274 3673.0     65   30121

Percentage of the requests served within a certain time (ms)
  50%     65
  66%    368
  75%    597
  80%    967
  90%   2953
  95%   7202
  98%  14507
  99%  21221
 100%  30121 (longest request)
```