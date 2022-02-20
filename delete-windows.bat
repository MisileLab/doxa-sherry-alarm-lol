SET startup-directory=%USERPROFILE%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup

if startup-directory/doxasherryalarm.lnk exists (
    del startup-directory/doxasherryalarm.lnk
) else (
    echo "설치되어있지 않은 것 같아요."
)