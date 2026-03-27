package i18n

import (
	"context"
)

type contextKey string

const languageContextKey contextKey = "language"

const (
	LangEN = "en"
	LangRU = "ru"
)

var DefaultLanguage = LangEN

var translations = map[string]map[string]string{
	LangEN: enTranslations,
	LangRU: ruTranslations,
}

func SetLanguage(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, languageContextKey, lang)
}

func GetLanguage(ctx context.Context) string {
	if lang, ok := ctx.Value(languageContextKey).(string); ok {
		return lang
	}
	return DefaultLanguage
}

func T(ctx context.Context, key string) string {
	lang := GetLanguage(ctx)
	if trans, ok := translations[lang]; ok {
		if val, ok := trans[key]; ok {
			return val
		}
	}
	// Fallback to English
	if trans, ok := translations[LangEN]; ok {
		if val, ok := trans[key]; ok {
			return val
		}
	}
	return key
}
