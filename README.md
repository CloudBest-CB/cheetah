# cheetah
基于open-falcon的开源监控系统

一、	程序下载
a)	在cloudbest官网、github、osChina下载程序的源码包或者编译好的二进制文件。
b)	将源码包进行编译前需要首先使用go get *** 命令安装项目依赖。
c)	安装所有依赖后使用go build命令在agent和server文件夹下编译。
二、	环境准备
a)	本监控器只支持linux环境，暂不支持windows和mac环境。
b)	如要使用源码进行编译安装，则需提前在系统安装go。
c)	本软件在go 1.10 版本调试通过，更低版本不解决兼容性问题。
三、	agent安装
a)	编译出二进制文件或直接下载
b)	将zkzd-agent与cfg.json和control文件放于系统内任意位置。
四、	agent启动
a)	在agent文件夹下的cfg.example.json文件，是配置文件cfg.json的示例，请根据具体配置修改其中的：server部分的addrs的值后将文件名改为cfg.json。
b)	将zkzd-agent与cfg.json和control文件放于文件夹后，使用chmod +x control zkzd-agent 命令
c)	最后执行 ./control start 启动
五、	server安装
a)	编译出二进制文件或直接下载
b)	将zkzd-server与cfg.json和control文件放于系统内任意位置。
六、	server启动
a)	在agent文件夹下的cfg.example.json文件，是配置文件cfg.json的示例，请根据具体配置修改其中的：influxAddr、database、username、password和myMeasurement部分的值后将文件名改为cfg.json。
b)	将zkzd-agent与cfg.json和control文件放于文件夹后，使用chmod +x control zkzd-agent 命令
c)	最后执行 ./control start 启动
七、	Influxdb安装
a)	在influxdb官网进行选择。https://portal.influxdata.com/downloads
b)	Ubuntu下：
Wget https://dl.influxdata.com/influxdb/nightlies/influxdb_nightly_amd64.deb
sudo dpkg -i influxdb_nightly_amd64.deb
c)	在centos下：
wget https://dl.influxdata.com/influxdb/nightlies/influxdb-nightly.x86_64.rpm
sudo yum localinstall influxdb-nightly.x86_64.rpm
八、	Influxdb配置与启动
a)	在安装之后，根据需要，在/etc/influx/influx.conf中更改监听端口及监听地址段等配置。
b)	完成配置文件更改后，使用service influxdb start 启动。
九、	Grafana安装
a)	在Grafana官网进行选择。https://grafana.com/grafana/download
b)	Ubuntu下：
wget https://s3-us-west-2.amazonaws.com/grafana-releases/release/grafana_5.2.2_amd64.deb 
sudo dpkg -i grafana_5.2.2_amd64.deb
c)	在centos下：
wget https://s3-us-west-2.amazonaws.com/grafana-releases/release/grafana-5.2.2-1.x86_64.rpm 
sudo yum localinstall grafana-5.2.2-1.x86_64.rpm
十、	Grafana配置与启动
a)	先进行启动服务service grafana-server start
b)	再进入页面http://127.0.0.1:3000登陆grafana的管理页面
c)	进行数据源配置，选择influxdb作为默认数据源，填入地址及用户名密码后保存
d)	创建新的展示大盘，根据需要新建图表。
