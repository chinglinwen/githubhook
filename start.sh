pkill hugo
pkill githubhook

./githubhook &> hook.log &

cd blogs/
hugo server --bind 0.0.0.0 -b http://g.clwen.com/ --appendPort=false -D &
# hugo server --bind 0.0.0.0 &
cd -