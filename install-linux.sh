#!/bin/sh

if test -f "/etc/systemd/system/DoxaSherryAlarm.service"; then
    echo "이미 설치되어 있는 것 같아요, 만약 처음 깔았거나 삭제 후 다시 깔았다면 뭔가 잘못된 것입니다."
else
    if test -f "$(pwd)/config.json"; then
        cat <<- EOM | sudo tee /etc/systemd/system/DoxaSherryAlarm.service
[Unit]
Description=Doxa Sherry Alarm

[Service]
ExecStart=$(pwd)/start
WorkingDirectory=$(pwd)

[Install]
WantedBy=multi-user.target
EOM
        sudo systemctl daemon-reload
        sudo systemctl start DoxaSherryAlarm
        sudo systemctl enable DoxaSherryAlarm
        echo "설치 완료!"
    else
        echo "최초 실행인 것 같아요. 한번 실행하고 설치를 시작해주세요."
        exit 1
    fi
fi
