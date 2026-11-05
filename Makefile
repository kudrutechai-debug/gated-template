.PHONY: check go-check py-check

## Run everything the GitHub gate runs. Do this before you push.
check: go-check py-check

go-check:
	$(MAKE) -C go check

py-check:
	$(MAKE) -C python check
