# mas2nm
**masscan to nmap**  
masscan 粗扫 + nmap 精确扫描工具

# 用途
nmap 和 masscan 都是我常用的工具，这两个各有优缺点  

| scanner | speed | accuracy |
| ------- | ----- | -------- |
|  nmap   |  low  |   high   |
| masscan |  high |    low   |  

所以现在业界的通用方案是先masscan粗扫， 再用 nmap 对扫出来的结果再细扫一遍  
各个资产表导来导去很麻烦，所以现在衍生了很多python脚本，但是python的各种依赖、版本问题让我苦不堪言  。

于是我想来用 golang 重写一遍，再静态编译打包好。
这样在任何一个服务器上、vps上都能快速开启扫描，这可就方便太多了。  
这便是这个项目的初衷

# TODO
- [x] masscan调用
- [x] masscan结果解析
- [ ] nmap调用
