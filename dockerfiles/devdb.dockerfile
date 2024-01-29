FROM postgres:15

COPY citext.sql /docker-entrypoint-initdb.d

ENV POSTGRES_USER=temaadmin
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=tema