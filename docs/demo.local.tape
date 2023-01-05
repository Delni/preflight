# Where should we write the GIF?
Output demo.local.gif

# Set up a 1200x600 terminal with 46px font.
Set FontSize 20
Set Width 1200
Set Height 600

# Type a command in the terminal.
Sleep 1s
Type "preflight --help"

# Pause for dramatic effect...
Sleep 500ms

# Run the command by pressing enter.
Enter

# Clean the preview for next command
Sleep 10s
Ctrl+L
Sleep 1s

# Type a command in the terminal.
Type "preflight docs/demo.yaml"
Sleep 500ms
Enter

# Clean the preview for next command

Sleep 10s
Ctrl+L
Sleep 1s

# Type a command in the terminal.
Type "preflight --checklists node,node-version,yarn"
Sleep 500ms
Enter

# Admire the output for a bit.
Sleep 10s
Ctrl+L