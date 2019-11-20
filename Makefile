OPENAPI_DIR=openapi

.PHONY: all openapi clean

all: openapi
 
openapi:
	@(cd $(OPENAPI_DIR) && $(MAKE))

clean:
	@(cd $(OPENAPI_DIR) && $(MAKE) $@)
