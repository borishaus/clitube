# TODO: Implement "play" Command

## Problem
Currently users run `clitube lofigirl` to play. This causes conflicts if someone names an alias "add", "rm", "list", etc.

## Goal
Change to: `clitube play lofigirl`

## Files to Modify
Only **main.go** needs changes!

## Steps to Implement

### 1. Add "play" case to switch statement (around line 221)
Look for the `switch command {` section. You need to:
- Add a new `case "play":`
- Call handlePlay with args, similar to how "add" and "remove" work
- Decide what to do with the `default:` case (maybe show an error for unknown commands?)

### 2. Update handlePlay function signature (around line 125)
Currently: `func handlePlay(alias string, videoMode bool)`

Consider changing to: `func handlePlay(args []string, videoMode bool)`
- Extract alias from `args[0]`
- Check that `len(args) >= 1` before accessing args[0]
- Return an error if no alias provided (like handleAdd and handleRemove do)

### 3. Update all help text
Search for TODO comments in:
- `printUsage()` function (line 8)
- `printFirstRunHints()` function (line 147)
- Update all examples from `clitube lofigirl` to `clitube play lofigirl`

### 4. Consider the -v flag behavior
Current: `clitube -v lofigirl`

Options after your change:
- Option A: `clitube -v play lofigirl` (flag before command)
- Option B: `clitube play -v lofigirl` (flag after command)

The current code structure checks for `-v` flag BEFORE reading the command (line 193).
So Option A would work without changes. For Option B, you'd need to check for `-v`
flag within the play command handler itself.

## Testing Your Changes

After implementing:
```bash
go build -o clitube

# Should work:
./clitube add myalias "https://youtube.com/watch?v=test"
./clitube play myalias
./clitube -v play myalias

# Should show error:
./clitube unknowncommand
./clitube play    # no alias provided
```

## Learning Resources
- Look at handleAdd() and handleRemove() - they show the pattern to follow
- The args variable is a "slice" (Go's version of an array)
- Use len(args) to check size, args[0] to get first element
- fmt.Errorf() creates error messages

Good luck! This is a great exercise for understanding:
- Command-line argument parsing
- Switch statements in Go
- Function signatures and parameters
- Error handling
