rd /s/q release
md release
go build -ldflags "-H windowsgui" -o area.exe
::go build -o area.exe
COPY area.exe release\
COPY favicon.ico release\favicon.ico
::XCOPY asset\*.* release\asset\  /s /e
XCOPY view\*.* release\view\  /s /e