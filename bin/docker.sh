#!/bin/sh
# 运行容器：docker run -d -v $prj_volume_host:$prj_volume_container -p $prj_port_host:$prj_port_container --name $prj_name $prj_image
# ＝》docker.sh --cmd=run --prj_image=leo-leogo --prj_name=leoGo --prj_port_host=8081 --prj_port_container=8081 --prj_volume_host=/data/leoGo --prj_volume_container=/data/leoGo
# 登录容器：docker exec -it [cid]|[cname] /bin/bash
# 查看容器文件卷：docker inspect -f "{{.Mounts}}" [cid]|[cname]
# docker宿主机相关文件目录说明
# 1、/var/lib/docker/devicemapper/devicemapper/data       #用来存储相关的存储池数据
# 2、/var/lib/docker/devicemapper/devicemapper/metadata   #用来存储相关的元数据。
# 3、/var/lib/docker/devicemapper/metadata/               #用来存储 device_id、大小、以及传输_id、初始化信息
# 4、/var/lib/docker/devicemapper/mnt                     #用来存储挂载信息
# 5、/var/lib/docker/container/                           #用来存储容器信息
# 6、/var/lib/docker/graph/                               #用来存储镜像中间件及本身详细信息和大小 、以及依赖信息
# 7、/var/lib/docker/repositores-devicemapper             #用来存储镜像基本信息
# 8、/var/lib/docker/tmp                                  #docker临时目录
# 9、/var/lib/docker/trust                                #docker信任目录
# 10、/var/lib/docker/volumes                             #docker卷目录
# 查看容器日志：docker logs --tail 0 -f [cid]|[cname]

# 项目容器名
#prj_name=leoGo
# 项目镜像地址
prj_image=leo-leogo
# 宿主机端口
#prj_port_host=8081
# 容器端口
#prj_port_container=8081

# initialize parameters
for arg in "$@"; do
  param=${arg%%=*}
  value=${arg#*=}
  case $param in
    --cmd)
      cmd=$value;;
    --prj_name)
      prj_name=$value;;
    --prj_image)
      prj_image=$value;;
    --prj_port_host)
      prj_port_host=$value;;
    --prj_port_container)
      prj_port_container=$value;;
    --prj_volume_host)
      prj_volume_host=$value;;
    --prj_volume_container)
      prj_volume_container=$value;;
    --help)
      echo "args:"
      echo "--cmd=run, login"
      echo "--prj_name="
      echo "--prj_image="
      echo "--prj_port_host="
      echo "--prj_port_container="
      echo "--prj_volume_host="
      echo "--prj_volume_container="
      echo "--help"
      exit 0;;
  esac
done

# verify parameters
if [ -z "$cmd" ] || [ -z "$prj_name" ]; then
    echo "parameters are invalid, please try --help"
    exit 1
fi

if [ "$cmd" == "run" ]; then
    if [ -z "$prj_image" ]; then
        echo "parameters are invalid, please try --help"
        exit 1
    fi

    # 服务器上删除任何没有用的image，使得服务器整洁并降低磁盘空间的占用.
    exec_ret=$(docker images --filter "dangling=true" -q --no-trunc)
    if [ -z "$exec_ret" ]; then
        echo ""
    else
        docker rmi $exec_ret
    fi
    # 运行前终止并移除退出的容器.
    docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker stop
    docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker rm
    docker images|grep none|awk '{print $3 }'|xargs docker rmi
    # 运行前终止并移除同名容器.
    if docker stop $prj_name; then docker rm $prj_name; fi

    # 基于prj_image镜像后台运行prj_name容器.
    docker run -d -v $prj_volume_host:$prj_volume_container -p $prj_port_host:$prj_port_container --name $prj_name $prj_image
elif [ "$cmd" == "login" ]; then
    docker exec -it $prj_name /bin/bash
fi
