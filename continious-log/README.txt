#Clone the repo
$cd continious-log
$make

then 

$ docker run continious-log-app:latest
2020/06/10 18:36:32 Test log to see on the console


Also you can run the container is dettach mode -

$ docker run -d continious-log-app:latest
45c0a0a71cc5a2362d838162e5e317945d97d7f98b8cb0e058734e33d16bacab

$ docker ps
CONTAINER ID  IMAGE                       COMMAND             CREATED             STATUS              PORTS   NAMES
45c0a0a71cc5  continious-log-app:latest   "continious-log"    4 seconds ago       Up 4 seconds                laughing_taussig


$ docker logs -f 45c0a0a71cc5
2020/06/10 18:45:12 Test log to see on the console
2020/06/10 18:45:22 Test log to see on the console
2020/06/10 18:45:32 Test log to see on the console
2020/06/10 18:45:42 Test log to see on the console
2020/06/10 18:45:52 Test log to see on the console
2020/06/10 18:46:02 Test log to see on the console
2020/06/10 18:46:12 Test log to see on the console
