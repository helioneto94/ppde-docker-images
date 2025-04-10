# # Como não ficou claro no enunciado como deveria ser passado para o container o video para ser convertido
# # eu optei por criar um diretório /videos onde inseri dois videos de teste. Esse diretorio é copiado para a imagem na hora do build
 
FROM golang:1.24.2 as build
WORKDIR /build
COPY convert.go .
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.21.3 as app
RUN apk add --no-cache ffmpeg
WORKDIR /app
COPY --from=build /build/main .
COPY ./videos .

ENTRYPOINT [ "./main" ]

# Deixei o CMD com dois argumentos em branco como padrão, para que quando o container for iniciado sem argumentos, o erro
# do próprio script apareça
# Deixei comentado um CMD padrão que poderia ser usado pra executar a conversão do video.mp4 para avi

# CMD ["video.mp4", "avi"]
CMD ["", ""]
