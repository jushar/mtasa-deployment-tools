package main

type ApiMessage struct {
	Status string `json:"status"`
}

type ConsoleOutputMessage struct {
	ApiMessage

	Output string `json:"output"`
}

type MTAStatusInfoMessage struct {
	Running        bool   `json:"running"`
	Process        string `json:"process"`
	ProcessProcess string `json:"process_process"`
	ProcessStatus  string `json:"process_status"`
}
