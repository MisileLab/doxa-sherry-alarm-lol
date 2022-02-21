SET working-directory=%cd%
SET startup-directory=%USERPROFILE%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup

if exist "%startup-directory%\doxasherryalarm.lnk" ( 
echo "Already installed now"
pause 
) else (
mklink "%startup-directory%\doxasherryalarm.lnk" "%working-directory%/start.exe"
copy "%working-directory%\config.json" "%startup-directory%"

echo "Installed"
pause
)