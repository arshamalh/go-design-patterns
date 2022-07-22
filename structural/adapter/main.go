package main

func main() {

	client := &Client{}

	// Compatible device //
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	// Non-compatible device //
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}