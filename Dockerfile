FROM "golang"

WORKDIR /app

RUN mkdir -p /usr/local/go/src/github.com/irwaters/alive bin
RUN groupadd appuser
RUN useradd -g appuser appuser
RUN chown -R appuser:appuser /app

ADD Makefile /app
ADD ./cmd /app/cmd
ADD ./pkg /usr/local/go/src/github.com/irwaters/alive/pkg

USER appuser

EXPOSE 8080/tcp

CMD ["make", "run"]

