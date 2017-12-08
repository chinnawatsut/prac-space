# Docker mySQL
open docker

docker pull mysql/mysql-server

#run dockker
docker run -p 3306:3306 --name mysql-container-practice -e MYSQL_ROOT_PASSWORD=rootroot -d mysql/mysql-server:latest

#connect to mysql as super user
docker exec -it mysql-container-practice mysql -uroot -p

#grant user
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'username'@'localhost' WITH GRANT OPTION;
CREATE USER 'username'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'username'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

#connect with workbench
username@0.0.0.0:3306 -p password

#list all docker stopped
docker ps -a

#stop docker
docker stop <tag>

#start docker
docker start <tag>