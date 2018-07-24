FROM "golang"

WORKDIR /app

RUN mkdir -p /usr/local/go/src/github.com/irwaters/truth_service
ADD Makefile /app
ADD ./cmd /app/cmd
ADD ./pkg /usr/local/go/src/github.com/irwaters/truth_service/pkg

EXPOSE 8080/tcp

CMD ["make", "run"]

