FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .main.go

EXPOSE 8080
# CMD [ "/app/main" ]

# FROM postgres:16

# COPY db/init.sql /docker-entrypoint-initdb.d/

# ENV POSTGRES_USER=root
# ENV POSTGRES_PASSWORD=password
# ENV POSTGRES_DB=todo

# CMD ["docker-entrypoint.sh", "postgres"]