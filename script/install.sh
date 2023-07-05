#!/bin/sh
git clone https://github.com/MapleLeafTeam/Hikari-Core.git
echo "下载core完成"
pip install poetry
echo "安装poetry完成"
cd Hikari-Core && poetry install
echo "安装依赖完成"
python init.py
echo "初始化完成"
poetry run uvicorn main:app --host 0.0.0.0 --port 8080
echo "它在运行了！"