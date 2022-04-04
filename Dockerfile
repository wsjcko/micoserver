FROM alpine
ADD micoserver /micoserver
ENTRYPOINT ["/micoserver"]