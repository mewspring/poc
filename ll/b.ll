global i32 42

define i32 @main() {
	%foo = load i32, i32* @0
	ret i32 %foo
}
