declare void @foo()
unnamed_addr global i32 42

; nothing interesting below. added to make lli happy.
define i32 @main() {
	ret i32 0
}
