# Lab0 实验报告

## 实验结果

### 1. 完成 Map-Reduce 框架

```
go test -v -timeout 60m -run=TestExampleURLTop
=== RUN   TestExampleURLTop
Case0 PASS, dataSize=1MB, nMapFiles=5, cost=58.9685ms
Case1 PASS, dataSize=1MB, nMapFiles=5, cost=35.5182ms
Case2 PASS, dataSize=1MB, nMapFiles=5, cost=37.9978ms
Case3 PASS, dataSize=1MB, nMapFiles=5, cost=47.624ms
Case4 PASS, dataSize=1MB, nMapFiles=5, cost=94.8831ms
Case5 PASS, dataSize=1MB, nMapFiles=5, cost=53.0009ms
Case6 PASS, dataSize=1MB, nMapFiles=5, cost=51.1397ms
Case7 PASS, dataSize=1MB, nMapFiles=5, cost=52.5168ms
Case8 PASS, dataSize=1MB, nMapFiles=5, cost=52.3723ms
Case9 PASS, dataSize=1MB, nMapFiles=5, cost=49.8744ms
Case10 PASS, dataSize=1MB, nMapFiles=5, cost=41.6575ms
Case0 PASS, dataSize=10MB, nMapFiles=10, cost=368.0996ms
Case1 PASS, dataSize=10MB, nMapFiles=10, cost=183.9994ms
Case2 PASS, dataSize=10MB, nMapFiles=10, cost=124.3802ms
Case3 PASS, dataSize=10MB, nMapFiles=10, cost=140.5764ms
Case4 PASS, dataSize=10MB, nMapFiles=10, cost=717.2154ms
Case5 PASS, dataSize=10MB, nMapFiles=10, cost=355.3363ms
Case6 PASS, dataSize=10MB, nMapFiles=10, cost=328.3177ms
Case7 PASS, dataSize=10MB, nMapFiles=10, cost=310.331ms
Case8 PASS, dataSize=10MB, nMapFiles=10, cost=238.6216ms
Case9 PASS, dataSize=10MB, nMapFiles=10, cost=204.8433ms
Case10 PASS, dataSize=10MB, nMapFiles=10, cost=164.9144ms
Case0 PASS, dataSize=100MB, nMapFiles=20, cost=4.2394058s
Case1 PASS, dataSize=100MB, nMapFiles=20, cost=1.8043909s
Case2 PASS, dataSize=100MB, nMapFiles=20, cost=1.5191705s
Case3 PASS, dataSize=100MB, nMapFiles=20, cost=1.3798997s
Case4 PASS, dataSize=100MB, nMapFiles=20, cost=3.4911209s
Case5 PASS, dataSize=100MB, nMapFiles=20, cost=3.2398985s
Case6 PASS, dataSize=100MB, nMapFiles=20, cost=2.9905075s
Case7 PASS, dataSize=100MB, nMapFiles=20, cost=4.9512063s
Case8 PASS, dataSize=100MB, nMapFiles=20, cost=3.5682317s
Case9 PASS, dataSize=100MB, nMapFiles=20, cost=4.268662s
Case10 PASS, dataSize=100MB, nMapFiles=20, cost=2.6018291s
Case0 PASS, dataSize=500MB, nMapFiles=40, cost=29.9564127s
Case1 PASS, dataSize=500MB, nMapFiles=40, cost=20.8204896s
Case2 PASS, dataSize=500MB, nMapFiles=40, cost=11.7497502s
Case3 PASS, dataSize=500MB, nMapFiles=40, cost=12.1077916s
Case4 PASS, dataSize=500MB, nMapFiles=40, cost=11.4297612s
Case5 PASS, dataSize=500MB, nMapFiles=40, cost=25.3385212s
Case6 PASS, dataSize=500MB, nMapFiles=40, cost=21.5388647s
Case7 PASS, dataSize=500MB, nMapFiles=40, cost=23.0925215s
Case8 PASS, dataSize=500MB, nMapFiles=40, cost=16.8624234s
Case9 PASS, dataSize=500MB, nMapFiles=40, cost=16.0360524s
Case10 PASS, dataSize=500MB, nMapFiles=40, cost=14.0815862s
Case0 PASS, dataSize=1GB, nMapFiles=60, cost=51.468262s
Case1 PASS, dataSize=1GB, nMapFiles=60, cost=1m11.8568274s
Case2 PASS, dataSize=1GB, nMapFiles=60, cost=44.0676332s
Case3 PASS, dataSize=1GB, nMapFiles=60, cost=35.0600089s
Case4 PASS, dataSize=1GB, nMapFiles=60, cost=23.3730395s
Case5 PASS, dataSize=1GB, nMapFiles=60, cost=50.1442693s
Case6 PASS, dataSize=1GB, nMapFiles=60, cost=1m0.8935218s
Case7 PASS, dataSize=1GB, nMapFiles=60, cost=1m0.5637238s
Case8 PASS, dataSize=1GB, nMapFiles=60, cost=39.9150889s
Case9 PASS, dataSize=1GB, nMapFiles=60, cost=2m26.3230141s
Case10 PASS, dataSize=1GB, nMapFiles=60, cost=2m9.8778425s
--- PASS: TestExampleURLTop (1259.44s)
PASS
ok     talent 1259.742s
```

result of test example is available in log/test_example.log

### 2. 基于 Map-Reduce 框架编写 Map-Reduce 函数

```
go test -v -timeout 60m -run=TestURLTop
=== RUN   TestURLTop
Case0 PASS, dataSize=1MB, nMapFiles=5, cost=18.353ms
Case1 PASS, dataSize=1MB, nMapFiles=5, cost=32.6882ms
Case2 PASS, dataSize=1MB, nMapFiles=5, cost=30.26ms
Case3 PASS, dataSize=1MB, nMapFiles=5, cost=37.8005ms
Case4 PASS, dataSize=1MB, nMapFiles=5, cost=77.9731ms
Case5 PASS, dataSize=1MB, nMapFiles=5, cost=26.0018ms
Case6 PASS, dataSize=1MB, nMapFiles=5, cost=23.9746ms
Case7 PASS, dataSize=1MB, nMapFiles=5, cost=22.9974ms
Case8 PASS, dataSize=1MB, nMapFiles=5, cost=27.0011ms
Case9 PASS, dataSize=1MB, nMapFiles=5, cost=28.2832ms
Case10 PASS, dataSize=1MB, nMapFiles=5, cost=24.0007ms
Case0 PASS, dataSize=10MB, nMapFiles=10, cost=44.9811ms
Case1 PASS, dataSize=10MB, nMapFiles=10, cost=49.4384ms
Case2 PASS, dataSize=10MB, nMapFiles=10, cost=53.9978ms
Case3 PASS, dataSize=10MB, nMapFiles=10, cost=81.9622ms
Case4 PASS, dataSize=10MB, nMapFiles=10, cost=438.4454ms
Case5 PASS, dataSize=10MB, nMapFiles=10, cost=44.0624ms
Case6 PASS, dataSize=10MB, nMapFiles=10, cost=45.7452ms
Case7 PASS, dataSize=10MB, nMapFiles=10, cost=44.5146ms
Case8 PASS, dataSize=10MB, nMapFiles=10, cost=59.991ms
Case9 PASS, dataSize=10MB, nMapFiles=10, cost=61.9243ms
Case10 PASS, dataSize=10MB, nMapFiles=10, cost=54.6125ms
Case0 PASS, dataSize=100MB, nMapFiles=20, cost=212.8366ms
Case1 PASS, dataSize=100MB, nMapFiles=20, cost=173.1011ms
Case2 PASS, dataSize=100MB, nMapFiles=20, cost=215.4799ms
Case3 PASS, dataSize=100MB, nMapFiles=20, cost=268.8863ms
Case4 PASS, dataSize=100MB, nMapFiles=20, cost=2.8741555s
Case5 PASS, dataSize=100MB, nMapFiles=20, cost=297.529ms
Case6 PASS, dataSize=100MB, nMapFiles=20, cost=348.6926ms
Case7 PASS, dataSize=100MB, nMapFiles=20, cost=3.2647504s
Case8 PASS, dataSize=100MB, nMapFiles=20, cost=482.773ms
Case9 PASS, dataSize=100MB, nMapFiles=20, cost=294.0718ms
Case10 PASS, dataSize=100MB, nMapFiles=20, cost=1.6772614s
Case0 PASS, dataSize=500MB, nMapFiles=40, cost=888.5318ms
Case1 PASS, dataSize=500MB, nMapFiles=40, cost=1.9729056s
Case2 PASS, dataSize=500MB, nMapFiles=40, cost=881.5675ms
Case3 PASS, dataSize=500MB, nMapFiles=40, cost=1.7402481s
Case4 PASS, dataSize=500MB, nMapFiles=40, cost=13.2728681s
Case5 PASS, dataSize=500MB, nMapFiles=40, cost=767.1916ms
Case6 PASS, dataSize=500MB, nMapFiles=40, cost=8.3611567s
Case7 PASS, dataSize=500MB, nMapFiles=40, cost=1.4271523s
Case8 PASS, dataSize=500MB, nMapFiles=40, cost=6.2572524s
Case9 PASS, dataSize=500MB, nMapFiles=40, cost=4.0495713s
Case10 PASS, dataSize=500MB, nMapFiles=40, cost=4.0049373s
Case0 PASS, dataSize=1GB, nMapFiles=60, cost=7.5740256s
Case1 PASS, dataSize=1GB, nMapFiles=60, cost=7.6084012s
Case2 PASS, dataSize=1GB, nMapFiles=60, cost=11.0803842s
Case3 PASS, dataSize=1GB, nMapFiles=60, cost=12.6859361s
Case4 PASS, dataSize=1GB, nMapFiles=60, cost=18.7231842s
Case5 PASS, dataSize=1GB, nMapFiles=60, cost=10.2685383s
Case6 PASS, dataSize=1GB, nMapFiles=60, cost=5.4361549s
Case7 PASS, dataSize=1GB, nMapFiles=60, cost=8.1474077s
Case8 PASS, dataSize=1GB, nMapFiles=60, cost=5.4648693s
Case9 PASS, dataSize=1GB, nMapFiles=60, cost=6.0677373s
Case10 PASS, dataSize=1GB, nMapFiles=60, cost=11.9352225s
--- PASS: TestURLTop (285.61s)
PASS
ok     talent 286.307s
```

result of test example is available in log/test_homework.log



## 实验总结

在这部分可以简单谈论自己在实验过程中遇到的困难、对 Map-Reduce 计算框架的理解，字数在 1000 字以内。

