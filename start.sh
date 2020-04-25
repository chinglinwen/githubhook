set -x
pkill githubhook
./githubhook &> hook.log &

cd blogs/
# hugo server --bind 0.0.0.0 -b http://g.clwen.com/ --appendPort=false &
./start.sh
cd -