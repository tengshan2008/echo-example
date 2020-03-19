server:
	go build -o /tmp/echo-example main.go
	nohup /tmp/echo-example>log.out 2>&1 &
stop:
	pkill echo-example