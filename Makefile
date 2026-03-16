DOC_OUT_DIR := ./doc
# 关键修改：增加 -not -path "./google/*"，不要手动去编译谷歌的标准库
ALL_PROTO_FILES := $(shell find . -name "*.proto" \
                    -not -path "./$(DOC_OUT_DIR)/*" \
                    -not -path "./google/*" \
                    -not -path "./temp_protos/*")

.PHONY: all openapi clean fix_package

all: openapi

fix_package:
	@echo "正在批量修正 go_package 设置..."
	@for file in $(ALL_PROTO_FILES); do \
	   if grep -q "option go_package" $$file; then \
	      sed -i '' 's|option go_package = "\(.*\)";|option go_package = "./\1";|g' $$file; \
	   else \
	      echo 'option go_package = "./pb";' >> $$file; \
	   fi \
	done

openapi:
	@echo "正在创建输出目录..."
	@mkdir -p $(DOC_OUT_DIR)
	@echo "正在生成 OpenAPI 文档..."
	# 注意：--proto_path=. 依然保留，用于 import 寻找，但 $(ALL_PROTO_FILES) 不再包含 google 里的文件
	protoc --proto_path=. \
    	       --openapi_out=version=3.0.0,naming=json:./doc \
    	       $(ALL_PROTO_FILES)
	@echo "------------------------------------------------"
	@echo "生成成功！"

clean:
	rm -rf $(DOC_OUT_DIR)