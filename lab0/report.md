# Lab0 实验报告

## 实验结果

### 1. 完成 Map-Reduce 框架

```
go test -v -timeout 60m -run=TestExampleURLTop
=== RUN   TestExampleURLTop
Case0 PASS, dataSize=1MB, nMapFiles=5, cost=54.1068ms
Case1 PASS, dataSize=1MB, nMapFiles=5, cost=35.8471ms
Case2 PASS, dataSize=1MB, nMapFiles=5, cost=34.1654ms
Case3 PASS, dataSize=1MB, nMapFiles=5, cost=43.7147ms
Case4 PASS, dataSize=1MB, nMapFiles=5, cost=79.54ms
Case5 PASS, dataSize=1MB, nMapFiles=5, cost=49.3617ms
Case6 PASS, dataSize=1MB, nMapFiles=5, cost=46.5659ms
Case7 PASS, dataSize=1MB, nMapFiles=5, cost=50.0676ms
Case8 PASS, dataSize=1MB, nMapFiles=5, cost=48.8755ms
Case9 PASS, dataSize=1MB, nMapFiles=5, cost=41.0224ms
Case10 PASS, dataSize=1MB, nMapFiles=5, cost=38.5016ms
Case0 PASS, dataSize=10MB, nMapFiles=10, cost=320.2324ms
Case1 PASS, dataSize=10MB, nMapFiles=10, cost=167.8428ms
Case2 PASS, dataSize=10MB, nMapFiles=10, cost=120.667ms
Case3 PASS, dataSize=10MB, nMapFiles=10, cost=119.6882ms
Case4 PASS, dataSize=10MB, nMapFiles=10, cost=641.7379ms
Case5 PASS, dataSize=10MB, nMapFiles=10, cost=321.3038ms
Case6 PASS, dataSize=10MB, nMapFiles=10, cost=301.4055ms
Case7 PASS, dataSize=10MB, nMapFiles=10, cost=364.9047ms
Case8 PASS, dataSize=10MB, nMapFiles=10, cost=223.7273ms
Case9 PASS, dataSize=10MB, nMapFiles=10, cost=178.8921ms
Case10 PASS, dataSize=10MB, nMapFiles=10, cost=158.2368ms
Case0 PASS, dataSize=100MB, nMapFiles=20, cost=3.1390011s
Case1 PASS, dataSize=100MB, nMapFiles=20, cost=1.3893513s
Case2 PASS, dataSize=100MB, nMapFiles=20, cost=954.1715ms
Case3 PASS, dataSize=100MB, nMapFiles=20, cost=2.4240299s
Case4 PASS, dataSize=100MB, nMapFiles=20, cost=2.7511003s
Case5 PASS, dataSize=100MB, nMapFiles=20, cost=3.6331775s
Case6 PASS, dataSize=100MB, nMapFiles=20, cost=3.2044724s
Case7 PASS, dataSize=100MB, nMapFiles=20, cost=3.9406203s
Case8 PASS, dataSize=100MB, nMapFiles=20, cost=2.0936492s
Case9 PASS, dataSize=100MB, nMapFiles=20, cost=1.5503635s
Case10 PASS, dataSize=100MB, nMapFiles=20, cost=1.2076629s
Case0 PASS, dataSize=500MB, nMapFiles=40, cost=25.4456317s
Case1 PASS, dataSize=500MB, nMapFiles=40, cost=14.8783234s
Case2 PASS, dataSize=500MB, nMapFiles=40, cost=10.4408874s
Case3 PASS, dataSize=500MB, nMapFiles=40, cost=10.4135559s
Case4 PASS, dataSize=500MB, nMapFiles=40, cost=10.8568517s
Case5 PASS, dataSize=500MB, nMapFiles=40, cost=21.310093s
Case6 PASS, dataSize=500MB, nMapFiles=40, cost=18.2229292s
Case7 PASS, dataSize=500MB, nMapFiles=40, cost=18.1044385s
Case8 PASS, dataSize=500MB, nMapFiles=40, cost=14.2397967s
Case9 PASS, dataSize=500MB, nMapFiles=40, cost=12.504795s
Case10 PASS, dataSize=500MB, nMapFiles=40, cost=14.0792216s
Case0 PASS, dataSize=1GB, nMapFiles=60, cost=48.314779s
Case1 PASS, dataSize=1GB, nMapFiles=60, cost=29.9605215s
Case2 PASS, dataSize=1GB, nMapFiles=60, cost=21.816549s
Case3 PASS, dataSize=1GB, nMapFiles=60, cost=20.0377209s
Case4 PASS, dataSize=1GB, nMapFiles=60, cost=20.6936616s
Case5 PASS, dataSize=1GB, nMapFiles=60, cost=47.7212088s
Case6 PASS, dataSize=1GB, nMapFiles=60, cost=41.5137968s
Case7 PASS, dataSize=1GB, nMapFiles=60, cost=44.1715308s
Case8 PASS, dataSize=1GB, nMapFiles=60, cost=32.9036035s
Case9 PASS, dataSize=1GB, nMapFiles=60, cost=24.7098543s
Case10 PASS, dataSize=1GB, nMapFiles=60, cost=2m32.8035366s
--- PASS: TestExampleURLTop (881.19s)
PASS
ok  	talent	881.572s
```

Result of "make test_example" is available in <u>log/test_example.log</u>



### 2. 基于 Map-Reduce 框架编写 Map-Reduce 函数

```
go test -v -timeout 60m -run=TestURLTop
=== RUN   TestURLTop
Case0 PASS, dataSize=1MB, nMapFiles=5, cost=25.0578ms
Case1 PASS, dataSize=1MB, nMapFiles=5, cost=29.6926ms
Case2 PASS, dataSize=1MB, nMapFiles=5, cost=282.7606ms
Case3 PASS, dataSize=1MB, nMapFiles=5, cost=77.0294ms
Case4 PASS, dataSize=1MB, nMapFiles=5, cost=38.2195ms
Case5 PASS, dataSize=1MB, nMapFiles=5, cost=185.0714ms
Case6 PASS, dataSize=1MB, nMapFiles=5, cost=30.1719ms
Case7 PASS, dataSize=1MB, nMapFiles=5, cost=30.1969ms
Case8 PASS, dataSize=1MB, nMapFiles=5, cost=94.7232ms
Case9 PASS, dataSize=1MB, nMapFiles=5, cost=28.062ms
Case10 PASS, dataSize=1MB, nMapFiles=5, cost=20.1835ms
Case0 PASS, dataSize=10MB, nMapFiles=10, cost=40.7678ms
Case1 PASS, dataSize=10MB, nMapFiles=10, cost=47.6344ms
Case2 PASS, dataSize=10MB, nMapFiles=10, cost=54.7339ms
Case3 PASS, dataSize=10MB, nMapFiles=10, cost=77.6767ms
Case4 PASS, dataSize=10MB, nMapFiles=10, cost=261.5726ms
Case5 PASS, dataSize=10MB, nMapFiles=10, cost=310.2633ms
Case6 PASS, dataSize=10MB, nMapFiles=10, cost=44.0939ms
Case7 PASS, dataSize=10MB, nMapFiles=10, cost=50.4649ms
Case8 PASS, dataSize=10MB, nMapFiles=10, cost=50.4096ms
Case9 PASS, dataSize=10MB, nMapFiles=10, cost=58.3175ms
Case10 PASS, dataSize=10MB, nMapFiles=10, cost=45.2759ms
Case0 PASS, dataSize=100MB, nMapFiles=20, cost=146.0478ms
Case1 PASS, dataSize=100MB, nMapFiles=20, cost=544.0832ms
Case2 PASS, dataSize=100MB, nMapFiles=20, cost=235.3012ms
Case3 PASS, dataSize=100MB, nMapFiles=20, cost=1.4182167s
Case4 PASS, dataSize=100MB, nMapFiles=20, cost=1.6650558s
Case5 PASS, dataSize=100MB, nMapFiles=20, cost=315.7017ms
Case6 PASS, dataSize=100MB, nMapFiles=20, cost=153.0375ms
Case7 PASS, dataSize=100MB, nMapFiles=20, cost=161.2142ms
Case8 PASS, dataSize=100MB, nMapFiles=20, cost=3.6473142s
Case9 PASS, dataSize=100MB, nMapFiles=20, cost=201.0136ms
Case10 PASS, dataSize=100MB, nMapFiles=20, cost=382.0689ms
Case0 PASS, dataSize=500MB, nMapFiles=40, cost=3.9562808s
Case1 PASS, dataSize=500MB, nMapFiles=40, cost=8.7393114s
Case2 PASS, dataSize=500MB, nMapFiles=40, cost=2.5093273s
Case3 PASS, dataSize=500MB, nMapFiles=40, cost=4.3728053s
Case4 PASS, dataSize=500MB, nMapFiles=40, cost=7.8807145s
Case5 PASS, dataSize=500MB, nMapFiles=40, cost=4.0481183s
Case6 PASS, dataSize=500MB, nMapFiles=40, cost=5.6082772s
Case7 PASS, dataSize=500MB, nMapFiles=40, cost=3.2234075s
Case8 PASS, dataSize=500MB, nMapFiles=40, cost=2.1543608s
Case9 PASS, dataSize=500MB, nMapFiles=40, cost=3.4617415s
Case10 PASS, dataSize=500MB, nMapFiles=40, cost=534.848ms
Case0 PASS, dataSize=1GB, nMapFiles=60, cost=9.7468331s
Case1 PASS, dataSize=1GB, nMapFiles=60, cost=8.9736291s
Case2 PASS, dataSize=1GB, nMapFiles=60, cost=7.1103251s
Case3 PASS, dataSize=1GB, nMapFiles=60, cost=12.0123006s
Case4 PASS, dataSize=1GB, nMapFiles=60, cost=17.6703582s
Case5 PASS, dataSize=1GB, nMapFiles=60, cost=6.1326302s
Case6 PASS, dataSize=1GB, nMapFiles=60, cost=8.0123164s
Case7 PASS, dataSize=1GB, nMapFiles=60, cost=9.4003026s
Case8 PASS, dataSize=1GB, nMapFiles=60, cost=6.8128708s
Case9 PASS, dataSize=1GB, nMapFiles=60, cost=5.1652438s
Case10 PASS, dataSize=1GB, nMapFiles=60, cost=9.4698508s
--- PASS: TestURLTop (275.26s)
PASS
ok  	talent	275.546s
```

Result of "make test_homework" is available in <u>log/test_homework.log</u>



## 实验总结

As a beginner of distributed system, I learned a lot from this lab. So, I have written a blog to summarize this lab, which is available in [https://blog.csdn.net/qq_36456827/article/details/119280942](https://blog.csdn.net/qq_36456827/article/details/119280942).

