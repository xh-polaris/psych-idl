# Proto 注解转换脚本

这个目录包含用于将 `google.api.http` 注解转换为 hz 工具的 `http` 注解格式的脚本。

## 背景

- **Swagger/OpenAPI 生成**：使用 `google.api.http` 注解（标准 gRPC-Gateway 格式）
- **Hz 代码生成**：使用 `http` 注解（CloudWeGo Hertz 自定义格式）

为了同时支持两种工具，我们在不同的生成阶段使用不同的注解格式。

## 脚本说明

### convert_to_hz_annotations.sh

将 `google.api.http` 注解转换为 hz 工具的 `http` 注解格式。

**使用方法：**
```bash
./scripts/convert_to_hz_annotations.sh <proto_file>
```

**转换示例：**
```protobuf
// 转换前
import "google/api/annotations.proto";
option (google.api.http) = {
  post: "/user/sign_in"
  body: "*"
};

// 转换后
import "basic/http.proto";
option (http.post) = "/user/sign_in";
```

## 在 Makefile 中的使用

```makefile
update-macos: convert-annotations
	hz --verbose update $(IDL_OPTIONS) --mod $(MODULE_NAME) $(EXTRA_OPTIONS)
	@find biz/application/dto -name "*.go" | xargs perl -i -pe 's/func init\(\).*//'
	@make restore-annotations
```

**流程：**
1. `convert-annotations`：将 proto 文件转换为 hz 格式
2. 运行 `hz update` 生成代码
3. `restore-annotations`：使用 `git checkout` 恢复 proto 文件为原始格式

## 注意事项

1. **IDL 分支**：始终保存使用 `google.api.http` 注解的原始 proto 文件
2. **代码生成**：临时转换为 hz 注解格式
3. **恢复**：生成完成后通过 git 恢复原始格式
4. **Git 提交**：确保 proto 文件的注解格式在 git 中保持一致

## 为什么使用 git 恢复而不是反向转换？

使用 `git checkout` 恢复原始文件而不是通过脚本反向转换的原因：

1. **格式保真**：git 恢复可以完全保留原始的缩进、空行和所有参数（如 `body: "*"`）
2. **简单可靠**：不需要维护复杂的双向转换逻辑
3. **避免误差**：脚本转换可能会丢失格式细节或引入错误
