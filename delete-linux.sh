#!/bin/sh

if test -f "/etc/systemd/system/DoxaSherryAlarm.service"; then
    sudo systemctl disable DoxaSherryAlarm
    sudo systemctl stop DoxaSherryAlarm
    sudo rm -rf /etc/systemd/system/DoxaSherryAlarm.service
    echo "시작프로그램 등록을 취소하였습니다."
else
    echo "시작프로그램으로 등록이 안되어 있는 것 같아요."
fi