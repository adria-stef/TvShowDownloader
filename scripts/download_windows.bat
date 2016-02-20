set path=%1
set name=%2
set tracker=%3

START cmd /c cd %path% && ctorrent -u %tracker% %name%
