# Kill Process Skill

## When to Use
When you need to kill a process but `pkill` is not available or user explicitly forbids it.

## Methods

### Method 1: ps + grep + kill (PID lookup)
```bash
ps aux | grep <process_name> | grep -v grep
# Find the PID (second column) from the output
kill <PID>
```

### Method 2: ps + grep + xargs kill (one-liner)
```bash
ps aux | grep <process_name> | grep -v grep | awk '{print $2}' | xargs kill
```

### Method 3: killall (if available)
```bash
killall <process_name>
```

### Method 4: Using pgrep + kill
```bash
kill $(pgrep -f <process_name>)
```

## Examples

### Kill emby-server
```bash
ps aux | grep emby-server | grep -v grep | awk '{print $2}' | xargs kill
```

### Kill process on specific port
```bash
lsof -ti:8096 | xargs kill
```

### Force kill if normal kill fails
```bash
ps aux | grep <process_name> | grep -v grep | awk '{print $2}' | xargs kill -9
```

## Important Notes
- Always use `grep -v grep` to exclude the grep process itself
- Use `kill -9` (SIGKILL) only as last resort - it doesn't allow graceful shutdown
- Check if process is gone after killing: `ps aux | grep <process_name>`
