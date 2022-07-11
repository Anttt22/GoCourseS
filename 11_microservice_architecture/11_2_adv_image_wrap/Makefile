REPOSITORY := rebrain
APP_NAME:=module11-task02
VERSION:=0.1

build:	
	@docker build -t $(REPOSITORY)/$(APP_NAME):$(VERSION) .	

run:		
	@docker run --name $(APP_NAME) --rm -d -p 8080:8080 $(REPOSITORY)/$(APP_NAME):$(VERSION)	

images:
	@docker images --format "{{.Repository}} {{.Size}}"

check:		
	@curl localhost:8080/

stop:
	@docker container stop $(APP_NAME)	
		  
