test工程下的PusherModuleTest的运行方法：
./PusherModuleTest <server> <session> <dir>


##1 编译步骤
###1 先编译BasicUsageEnvironment 
- 参考docs下的make_basicenv.so.txt
- 这个编译出来的动态库

###2 编译静态库libRTSPPusher.a
  - 使用Makefile.macosx 这个来
  - make -f Makefile.macosx all
###3 编译测试程序
   - 1 链接BasicUsageEnvironment
   - 2 链接rtsppusher库
   
   - 3 make -f Makefile.macosx test
   - 4 运行测试 

##2 运维
### 8554 占用

```
查看某个进程打开的文件句柄
sudo lsof -i -n -p 27691

查看端口号8080占用情况
sudo lsof -i -n -P | grep 8080
```
```
sudo lsof -i -n -P | grep 8554
Password:
live555Me 5655       zhangbin    3u  IPv4 0xdaa50a0010d9f17f      0t0    TCP *:8554 (LISTEN)
 zhangbin@bogon
```