repo:=awrenn53/relay-template-renderer
tag:=latest

image: 
	docker build -t $(repo):$(tag) .

upload: image
	docker push $(repo):$(tag)

run: image
	docker run -it $(repo):$(tag)
