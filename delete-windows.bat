SET startup-directory=%USERPROFILE%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup

if EXIST "%startup-directory%\doxasherryalarm.lnk" (
    del "%startup-directory%\doxasherryalarm.lnk"
    del "%startup-directory%\config.json"
) ELSE (
    echo "Not installed"
)

pause