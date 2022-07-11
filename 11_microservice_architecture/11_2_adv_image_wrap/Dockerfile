#собираем два контейнера - название техники - multistage
#первый контейнер -builder
FROM golang:latest as builder  
# builder - даем ему имя
# на какой опер системе собрать.
WORKDIR /app
# если нет такой папки то она создаст папаку и перейдет в нее
COPY . .
# копируем все файлы на хост машине на докер машину в папаку опис выше
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./rebrainapp ./main.go
# !!!what cases to goos t be changed to windows??? 
# компилируем приложение
# отключаем с 
# первый контейнер служит сборщиком

#контейнер scratch - маленький контенер (файловая система с возм запуска бинарников) 
# стартуем контейнер - он и будет результатом выполнения
FROM scratch
COPY --from=builder /app/rebrainapp /usr/bin/rebrainapp
#копируем из билдера
ENTRYPOINT ["/usr/bin/rebrainapp"]
# то что будет запущенно когда напишем докер ран