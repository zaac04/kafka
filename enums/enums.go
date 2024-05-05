package enums

const (
	Terminalreset     = "\033[0m"
	Terminalbold      = "\033[1m"
	Terminalunderline = "\033[4m"
	Terminalstrike    = "\033[9m"
	Terminalitalic    = "\033[3m"

	TerminalcRed    = "\033[31m"
	TerminalcGreen  = "\033[32m"
	TerminalcYellow = "\033[33m"
	TerminalcBlue   = "\033[34m"
	TerminalcPurple = "\033[35m"
	TerminalcCyan   = "\033[36m"
	TerminalcWhite  = "\033[37m"
)

const (
	EnvLoadFailed      = "Failed to load env"
	RedisJsonSetFailed = "Failed to set json in redis"
	RedisListSetFailed = "Failed to set json in redis"
)

const (
	NeverExpire = 0
)
