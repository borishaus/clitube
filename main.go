package main

import (
	"fmt"
	"os"
)

func printUsage() {
	// TODO: UPDATE HELP TEXT
	// After you implement the "play" command, you'll need to update these instructions!
	// Change "clitube <alias>" to "clitube play <alias>"
	// Also update "clitube -v <alias>" to "clitube -v play <alias>" or "clitube play -v <alias>"
	// depending on how you want the -v flag to work with the play command

	fmt.Println("cliTube - CLI YouTube Player")
	fmt.Println("\nStream YouTube videos using memorable aliases")
	fmt.Println("\nUsage:")
	fmt.Println("  clitube add <alias> <url>     Add a new video alias")
	fmt.Println("  clitube <alias>               Play audio from saved alias (default)")
	fmt.Println("  clitube -v <alias>            Play video from saved alias")
	fmt.Println("  clitube --video <alias>       Play video from saved alias")
	fmt.Println("  clitube list                  List all saved aliases")
	fmt.Println("  clitube remove <alias>        Remove an alias")
	fmt.Println("  clitube rm <alias>            Remove an alias (short form)")
	fmt.Println("  clitube help                  Show this help message")
	fmt.Println("\nExamples:")
	// TODO: UPDATE THESE EXAMPLES TOO!
	fmt.Println("  # Add a lofi music stream")
	fmt.Println("  clitube add lofigirl \"https://www.youtube.com/watch?v=jfKfPfyJRdk\"")
	fmt.Println("")
	fmt.Println("  # Play audio only (saves bandwidth)")
	fmt.Println("  clitube lofigirl")  // TODO: Change to "clitube play lofigirl"
	fmt.Println("")
	fmt.Println("  # Play with video")
	fmt.Println("  clitube -v lofigirl")  // TODO: Update this based on how you handle -v flag with play command
	fmt.Println("")
	fmt.Println("  # See what you have saved")
	fmt.Println("  clitube list")
	fmt.Println("")
	fmt.Println("  # Remove an alias")
	fmt.Println("  clitube rm lofigirl")
	fmt.Println("\nConfiguration:")
	fmt.Println("  Aliases are stored in: ~/.config/clitube/videos.json")
	fmt.Println("\nDependencies:")
	fmt.Println("  Requires mpv to be installed (https://mpv.io)")
	fmt.Println("\nDocumentation:")
	fmt.Println("  man clitube                   View full manual page")
}

func handleAdd(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: clitube add <alias> <url>")
	}

	alias := args[0]
	url := args[1]

	if err := AddMapping(alias, url); err != nil {
		return fmt.Errorf("failed to add mapping: %w", err)
	}

	fmt.Printf("Successfully added alias '%s' for URL: %s\n", alias, url)
	return nil
}

func handleList() error {
	mappings, err := LoadMappings()
	if err != nil {
		return fmt.Errorf("failed to load mappings: %w", err)
	}

	// Show recently played first
	recent, err := GetRecentHistory()
	if err == nil && len(recent) > 0 {
		fmt.Println("Recently played:")
		for i, entry := range recent {
			mode := "audio"
			if entry.VideoMode {
				mode = "video"
			}
			fmt.Printf("  %d. %s (%s) - %s\n", i+1, entry.Alias, mode, entry.PlayedAt.Format("Jan 2 15:04"))
		}
		fmt.Println()
	}

	if len(mappings.Aliases) == 0 {
		fmt.Println("No aliases saved yet.")
		fmt.Println("\nAdd one with: clitube add <alias> <url>")
		return nil
	}

	fmt.Println("Saved aliases:")
	for alias, url := range mappings.Aliases {
		fmt.Printf("  %s -> %s\n", alias, url)
	}

	return nil
}

func handleRemove(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: clitube remove <alias>")
	}

	alias := args[0]

	mappings, err := LoadMappings()
	if err != nil {
		return fmt.Errorf("failed to load mappings: %w", err)
	}

	if _, exists := mappings.Aliases[alias]; !exists {
		return fmt.Errorf("alias '%s' not found", alias)
	}

	delete(mappings.Aliases, alias)

	if err := SaveMappings(mappings); err != nil {
		return fmt.Errorf("failed to save mappings: %w", err)
	}

	fmt.Printf("Successfully removed alias '%s'\n", alias)
	return nil
}

// TODO: UPDATE THIS FUNCTION SIGNATURE
// Currently: handlePlay(alias string, videoMode bool)
// After fix: You might want to change this to handlePlay(args []string, videoMode bool)
//            Then extract the alias from args[0], similar to handleAdd() and handleRemove()
// Don't forget to check if args has at least 1 element before accessing args[0]!

func handlePlay(alias string, videoMode bool) error {
	url, err := GetURL(alias)
	if err != nil {
		return err
	}

	// Track in history
	if err := AddToHistory(alias, url, videoMode); err != nil {
		// Don't fail playback if history fails, just log
		fmt.Fprintf(os.Stderr, "Warning: failed to save to history: %v\n", err)
	}

	return Play(url, videoMode)
}

func printFirstRunHints() {
	// TODO: UPDATE FIRST RUN HINTS TOO!
	fmt.Println("┌─────────────────────────────────────────────────────────────────┐")
	fmt.Println("│ Welcome to cliTube! Here are some tips to get started:         │")
	fmt.Println("└─────────────────────────────────────────────────────────────────┘")
	fmt.Println()
	fmt.Println("📝 Quick Start:")
	fmt.Println("   1. Add a video:    clitube add lofi \"https://youtube.com/watch?v=...\"")
	fmt.Println("   2. Play audio:     clitube lofi")        // TODO: Change to "clitube play lofi"
	fmt.Println("   3. Play video:     clitube -v lofi")     // TODO: Update based on your -v flag implementation
	fmt.Println()
	fmt.Println("💡 Tips:")
	fmt.Println("   • By default, only audio is streamed (saves bandwidth)")
	fmt.Println("   • Use -v or --video flag to stream video too")
	fmt.Println("   • Your last 3 played items are tracked and shown with 'list'")
	fmt.Println("   • Aliases are stored in ~/.config/clitube/videos.json")
	fmt.Println()
	fmt.Println("📚 For more help:")
	fmt.Println("   clitube help       Show usage examples")
	fmt.Println("   man clitube        View full manual (if installed)")
	fmt.Println()
}

func showRecentHistory() {
	recent, err := GetRecentHistory()
	if err != nil || len(recent) == 0 {
		return
	}

	fmt.Println("\n🕐 Recently played:")
	for i, entry := range recent {
		mode := "audio"
		if entry.VideoMode {
			mode = "video"
		}
		fmt.Printf("   %d. %s (%s)\n", i+1, entry.Alias, mode)
	}
	fmt.Println()
}

func main() {
	// Check for first run and show hints
	if firstRun, err := IsFirstRun(); err == nil && firstRun {
		printFirstRunHints()
	}

	if len(os.Args) < 2 {
		// Show recent history if available
		showRecentHistory()
		printUsage()
		os.Exit(1)
	}

	// Check for video flag
	videoMode := false
	startIdx := 1

	if os.Args[1] == "-v" || os.Args[1] == "--video" {
		videoMode = true
		startIdx = 2

		if len(os.Args) < 3 {
			fmt.Println("Error: -v flag requires an alias")
			printUsage()
			os.Exit(1)
		}
	}

	command := os.Args[startIdx]
	args := os.Args[startIdx+1:]

	var err error

	// TODO: FIX COMMAND STRUCTURE
	// Current problem: The "default" case assumes any unknown command is an alias to play.
	// This means if a user creates an alias called "add" or "rm", it conflicts with built-in commands!
	//
	// HINT: You want users to run "clitube play lofigirl" instead of "clitube lofigirl"
	// This means:
	// 1. Add a new case for "play" in the switch statement below
	// 2. The "play" case should expect an alias in the args (similar to how "add" works)
	// 3. Remove or change the "default" case since we don't want to assume unknown commands are aliases
	// 4. Think about what error message to show if someone runs "clitube unknowncommand"
	//
	// Look at how handleAdd() receives args - handlePlay() should work similarly!
	// Currently handlePlay(command, videoMode) receives the alias as the first parameter,
	// but after your fix it should receive it from args[0] instead.

	switch command {
	case "add":
		err = handleAdd(args)
	case "list":
		err = handleList()
	case "remove", "rm":
		err = handleRemove(args)
	case "help", "--help", "-h":
		printUsage()
		return
	default:
		// PROBLEM: This assumes any unknown command is an alias to play
		// This causes conflicts if someone names their video "add", "rm", "list", etc.
		err = handlePlay(command, videoMode)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
