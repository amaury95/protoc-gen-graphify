commit:
	@git commit -am "Release $(version)"

tag:
	@git tag $(version)

push:
	@git push origin master $(version)

release: commit tag push

install: release
	@go install github.com/amaury95/protoc-gen-go-tag@$(version)

