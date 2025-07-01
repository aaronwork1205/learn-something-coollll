FROM golang:latest

LABEL developer="aaronwork1205@outlook.com"
LABEL version="v1.0"

WORKDIR /model
COPY . /model

RUN chmod +x script.sh     # 确保脚本可执行

ENTRYPOINT ["./script.sh"]
CMD ["go", "run", "api.go"]
