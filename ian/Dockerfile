FROM mysql

ENV MYSQL_ROOT_PASSWORD admin
ENV MYSQL_DATABASE hackday

COPY ./sql/schema.sql /docker-entrypoint-initdb.d/schema.sql

EXPOSE 3306
