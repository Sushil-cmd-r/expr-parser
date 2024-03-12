run: clean
	@go build -o main; ./main

clean:
	@rm -f ./main
