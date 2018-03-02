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
