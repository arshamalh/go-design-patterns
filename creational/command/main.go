package main

func main() {
	tv := &Tv{}

	// Register commands
	onCmd := &OnCommand{
		device: tv,
	}
	offCmd := &OffCommand{
		device: tv,
	}

	// Make buttons
	onBtn := &Button{
		command: onCmd,
	}

	offBtn := &Button{
		command: offCmd,
	}

	// Press buttons!
	onBtn.press()
	offBtn.press()
}