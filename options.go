package whatlanggo

// Options represents options that can be set when detecting a language or/and script such
// blacklisting languages to skip checking.

//把需要检测的语言放白的,不需要检测的放黑的. 默认不写就是所有都检测.
type Options struct {
	Whitelist map[Lang]bool
	Blacklist map[Lang]bool
}
