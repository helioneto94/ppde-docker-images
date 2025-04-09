
# Como não ficou claro no enunciado como deveria ser passado para o container o video para ser convertido
# eu optei por criar um diretório /videos onde inseri dois videos de teste. Esse diretorio é copiado para a imagem na hora do build
 
FROM ubuntu:22.04
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ffmpeg
WORKDIR /app
COPY convert.sh .
COPY ./videos .
RUN chmod +x convert.sh

ENTRYPOINT [ "./convert.sh" ]

# Deixei o CMD com dois argumentos em branco como padrão, para que quando o container for iniciado sem argumentos, o erro
# do próprio script apareça
# Deixei comentado um CMD padrão que poderia ser usado pra executar a conversão do video.mp4 para avi

# CMD ["video.mp4", "avi"]
CMD ["", ""]