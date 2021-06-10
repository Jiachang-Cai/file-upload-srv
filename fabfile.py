# -*- coding: utf-8 -*-
"""
fabric 部署脚本
pip install fabric3
fab build
fab deploy
"""

from fabric.api import *
import datetime

# 项目名称
PROJECT_NAME = "fmtsmsapi"
LOCAL_FILE_NAME = PROJECT_NAME
WORK_SPACE = "/data/workspace"
# 当天日期
CURRENT_DAY_STRING = datetime.datetime.today().strftime("%Y%m%d%H%M")

env.user = 'root'

def build():
    local("go-bindata  -pkg config -o config/bindata.go config/", capture=False)
    command = "GOOS=linux go build"
    local(command, capture=False)

@hosts("129.211.42.179")
def dev():
    """测试环境部署"""
    remote_run_name = WORK_SPACE + "/" + PROJECT_NAME + "/" + LOCAL_FILE_NAME
    remote_backup_name = WORK_SPACE + "/" + PROJECT_NAME + "/backup/" +  LOCAL_FILE_NAME + "_" + CURRENT_DAY_STRING
    # 创建远程文件目录
    back_up = WORK_SPACE + "/" + PROJECT_NAME + "/backup"
    run('mkdir -p %s' % back_up)

    # 将代码归档上传到服务器当中的临时文件夹内
    put(LOCAL_FILE_NAME, remote_backup_name)
    run("chmod +x %s" % remote_backup_name)
    # 替换正在执行得文件
    run("cp -rf %s %s" % (remote_backup_name, remote_run_name))
    # 重启服务
    run("sudo supervisorctl restart %s" % PROJECT_NAME)
    command = "rm %s" % PROJECT_NAME
    local(command, capture=False)

@hosts("172.81.237.24")
def deploy():
    """正式环境部署"""
    remote_run_name = WORK_SPACE + "/" + PROJECT_NAME + "/" + LOCAL_FILE_NAME
    remote_backup_name = WORK_SPACE + "/" + PROJECT_NAME + "/backup/" +  LOCAL_FILE_NAME + "_" + CURRENT_DAY_STRING
    # 创建远程文件目录
    back_up = WORK_SPACE + "/" + PROJECT_NAME + "/backup"
    run('mkdir -p %s' % back_up)

    # 将代码归档上传到服务器当中的临时文件夹内
    put(LOCAL_FILE_NAME, remote_backup_name)
    run("chmod +x %s" % remote_backup_name)
    # 替换正在执行得文件
    run("cp -rf %s %s" % (remote_backup_name, remote_run_name))
    # 重启服务
    run("sudo supervisorctl restart %s" % PROJECT_NAME)
    command = "rm %s" % PROJECT_NAME
    local(command, capture=False)