# Where should we write the GIF?
Output demo.remote.gif

# Set up a 1200x600 terminal with 46px font.
Set FontSize 20
Set Width 1200
Set Height 600

# Type a command in the terminal.
Sleep 1s
Type "preflight --remote "
Sleep 250ms
Type@10ms "https://raw.githubusercontent.com/delni/preflight/main/docs/demo.yaml"
Sleep 500ms
Enter

# Admire the output for a bit.
Sleep 10s
Ctrl+L