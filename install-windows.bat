SET working-directory=%cd%
SET startup-directory=%USERPROFILE%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup

if startup-directory/doxasherryalarm.lnk exists goto doxasherryalarm-exists

mklink startup-directory/doxasherryalarm.lnk working-directory/start.exe

chmod +x install-windows.bat

./install-windows.bat

echo "다 깔린 것 같아요."
exit 0

:doxasherryalarm-exists

echo "아마도 이미 깔린 것 같아요."