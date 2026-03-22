#!/bin/bash

# 将 proto 文件中的 google.api.http 注解转换为 hz 工具的 http 注解格式
# 使用方法：./scripts/convert_to_hz_annotations.sh <proto_file>

set -e

PROTO_FILE="$1"

if [ ! -f "$PROTO_FILE" ]; then
    echo "Error: File not found: $PROTO_FILE"
    exit 1
fi

echo "Converting annotations in: $PROTO_FILE"

# 使用 Python 脚本进行转换
python3 - "$PROTO_FILE" << 'PYTHON_SCRIPT'
import re
import sys

proto_file = sys.argv[1]

with open(proto_file, 'r') as f:
    content = f.read()

# 替换 import
content = content.replace('import "google/api/annotations.proto";', 'import "basic/http.proto";')

# 替换 google.api.http 注解块
def replace_http_annotation(match):
    annotation_body = match.group(1)
    
    # 提取 HTTP 方法和路径
    methods = ['post', 'get', 'put', 'patch', 'delete', 'head', 'options', 'any']
    for method in methods:
        pattern = rf'{method}:\s*"([^"]+)"'
        m = re.search(pattern, annotation_body)
        if m:
            path = m.group(1)
            return f'    option (http.{method}) = "{path}";'
    
    return match.group(0)

# 匹配多行注解块（包括 body 参数）
pattern = r'option\s+\(google\.api\.http\)\s*=\s*\{([^}]+)\};'
content = re.sub(pattern, replace_http_annotation, content, flags=re.DOTALL)

with open(proto_file, 'w') as f:
    f.write(content)

print(f"Converted: {proto_file}")
PYTHON_SCRIPT

echo "Done!"
