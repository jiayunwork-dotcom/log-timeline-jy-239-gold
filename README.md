# log-timeline

多源日志关联分析与故障时间线重建（CLI）。

将多行日志（可来自多个来源）解析为带时间戳的条目，按时间排序重建统一时间线，
并检测相邻条目之间的异常时间空隙（gap）。

## 用法

```
log-timeline -input logs.txt [-gap 5m] [-out -]
```

- `-input`  日志文件，`-` 表示标准输入；每行格式 `2006-01-02T15:04:05Z07:00 LEVEL [source] message`
- `-gap`    空隙阈值（Go duration，如 `5m`、`30s`）；相邻间隔 ≥ 阈值记为一次 gap
- `-out`    时间线输出，`-` 表示标准输出

## 行格式

```
2026-01-01T10:00:00Z INFO [api] request received
```

时间戳为 RFC3339；`LEVEL` 大小写不限；`[source]` 可选；其余为 message。
