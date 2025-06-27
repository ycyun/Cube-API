package utils

import (
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

//
//func (ch *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
//	return ch.Handler.Enabled(ctx, level)
//}
//
//func (ch *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
//	r.AddAttrs(getContextData(ctx)...)
//	return ch.Handler.Handle(ctx, r)
//}
//
//func (ch *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
//	println(attrs)
//	return &ContextHandler{ch.Handler.WithAttrs(attrs)}
//}
//
//func (ch *ContextHandler) WithGroup(name string) slog.Handler {
//	return &ContextHandler{ch.Handler.WithGroup(name)}
//}
//
//func getContextData(ctx context.Context) []slog.Attr {
//	var as []slog.Attr
//
//	// 전달받은 Context에서 host와 method를 꺼내서 slog.group 추가함.
//	group := slog.Group(
//		"request",
//		slog.String("host", getStringValue(ctx, "host")),
//		slog.String("method", getStringValue(ctx, "method")),
//	)
//	as = append(as, group)
//	return as
//}
//
//func getStringValue(ctx context.Context, key string) string {
//	value := ""
//	ctxValue := ctx.Value(key)
//	if ctxValue != nil {
//		stringValue, ok := ctxValue.(string)
//		if !ok {
//			return value
//		}
//		value = stringValue
//	}
//	return value
//}
