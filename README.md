# Doxa-Sherry-Alarm

[독사](https://www.twitch.tv/doxa_97)와 [쉐리](https://www.twitch.tv/20_sherry_02)의 음성으로 만들어진 엄청난 알람

## 설치

[MacOS](https://nightly.link/MisileLab/doxa-sherry-alarm-lol/workflows/build-macos/main/build-artifacts.zip)

### 윈도우

1. [여기서](https://nightly.link/MisileLab/doxa-sherry-alarm-lol/workflows/build-windows/main/build-artifacts.zip) 다운로드해주세요.
2. **한번 실행해주세요.**
3. install-windows.bat를 관리자 권한으로 실행시켜주세요.

### 리눅스

1. [여기서](https://nightly.link/MisileLab/doxa-sherry-alarm-lol/workflows/build/main/build-artifacts.zip) 다운로드해주세요.
2. **한번 실행해주세요.**
3. install-linux.sh를 sudo 권한이 있는 상태로 실행시켜주세요.

## 삭제

### 윈도우

delete-windows.bat를 관리자 권한이 있는 상태로 실행시켜주세요.

### 리눅스

delete-linux.sh를 관리자 권한이 있는 상태로 실행시켜주세요.

## 프로그램 설정 초기화

혹시 자는 시간을 잘못 설정했거나 일어나는 시간이 변경되었는데 고치고 싶으시다면 두가지 방법이 있습니다.

config.json을 삭제하거나,
프로그램 앞에 reset=true를 붙이시면 됩니다.
예) ```reset=true ./start.exe```
