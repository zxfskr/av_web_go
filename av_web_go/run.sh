# if [ "$1" = 'stop' ];then
#     echo "killed"
#     kill -9 `cat log/pid`
# else
#     echo "run"
#     nohup ./bin/av_web >/dev/null 2>&1 & echo $! > log/pid
# fi
./bin/av_web
