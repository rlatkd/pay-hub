FROM mysql:8.0
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_USER=rlatkd
ENV MYSQL_PASSWORD='0000'
ENV MYSQL_DATABASE=CHANNEL
CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci", "--skip-character-set-client-handshake"]
COPY ./init.sql /docker-entrypoint-initdb.d/
