#!/bin/bash
WHITELIST=true
IP="192.168.20.141"


DIRECTORY=`dirname "$ABSOLUTE_FILENAME"`
echo 'Создание и старт контейнеров'
sudo docker network create net
sudo docker-compose -f $DIRECTORY/pg_dump/docker-compose.yml build || exit
sudo docker-compose -f $DIRECTORY/pg_dump/docker-compose.yml up --no-start || exit
sudo docker-compose -f $DIRECTORY/pg_dump/docker-compose.yml start || exit

echo 'Ждем 5 сек'
sleep 5

echo 'Создаем базы данных'
sudo su - postgres <<EOF
    psql -U postgres -h $IP -p 5432
    DROP DATABASE authors;
    drop database books;
    drop database publishers;
    drop database users;

    create database authors;
    create database books;
    create database publishers;
    create database users;
    \q
    exit
EOF
echo 'Создание таблиц и заполнение данными'
export PGPASSWORD="password"
psql -h $IP -p 5432 --username=postgres authors < $DIRECTORY/pg_dump/authors.sql || exit
psql -h $IP -p 5432 --username=postgres books < $DIRECTORY/pg_dump/books.sql || exit
psql -h $IP -p 5432 --username=postgres publishers < $DIRECTORY/pg_dump/publishers.sql || exit
psql -h $IP -p 5432 --username=postgres users < $DIRECTORY/pg_dump/users.sql || exit


#Запуск
echo "Запуск сервисов"
sudo docker-compose -f $DIRECTORY/docker-compose.yml build
sudo docker-compose -f $DIRECTORY/docker-compose.yml up



echo 'Остановка'
#остановка всех контейнеров
sudo docker stop $(sudo docker ps -aq)

#Удаление всех данных docker
if $WHITELIST; then
  echo "Удаление всех остановленных контейнеров Docker"
  sudo docker rm $(sudo docker ps -aq)
  echo "Удаление множества пользовательских данных Docker"
  sudo docker system prune
  echo "Удаление неиспользуемых сетей Docker"
  sudo docker network prune
  echo "Удаление всех пользовательских образов Docker"
  sudo docker rmi $(docker images -a -q)
  echo "Удаление всех пользовательских томов Docker"
  sudo docker volume prune
fi