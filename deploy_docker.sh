!/bin/bash
#项目容器名
prj_name=leoGo
#项目镜像地址
prj_image
#宿主机端口
prj_port_host=8081
#容器端口
prj_port_container＝8081

#运行前终止并移除同名容器.
if docker stop $prj_name; then docker rm $prj_name; fi
#基于prj_image镜像后台运行prj_name容器.
docker run -d -p $prj_port_host:$prj_port_container --name $prj_name $prj_image
#服务器上删除任何没有用的image,持服务器整洁并降低磁盘空间的占用.
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
