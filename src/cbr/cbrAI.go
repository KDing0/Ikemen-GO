package cbr

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

//All data relevant for the AI to operate
//See CBRData.proto and CBRRawFrames.proto for the structure of the data
type AIData struct {
	cbrData	*CBRData				//Data storage after the AI processed a replay
	rawFrames *CBRRawFrames 		//Data storage of a replay before processing
	recording bool
	replaying bool
	replayFrames []*CBRData_Frame	//Frames that the AI is ready to send over to the game for replaying
	replayIndex int
}
var aiData = AIData{
	cbrData: &CBRData{},
	rawFrames: &CBRRawFrames{},
	recording: false,
	replaying: false,
	replayIndex: 0,
}

//Naming for inputs, individual bits represent pressed buttons 0001 = 1 = up, 0011 = 3 = up + down
const (
	IB_PU int32 = 1 << iota
	IB_PD
	IB_PL
	IB_PR
	IB_A
	IB_B
	IB_C
	IB_X
	IB_Y
	IB_Z
	IB_S
	IB_D
	IB_W
	IB_M
	IB_anybutton = IB_A | IB_B | IB_C | IB_X | IB_Y | IB_Z | IB_S | IB_D | IB_W | IB_M
)

///Interface functions for the outside program-----------
func ToggleRecording() bool {
	if aiData.recording == true {
		//AiData.rawFrames.clearRecording()
		aiData.EndRecording()
	}else{
		aiData.InitializeRecording()
		aiData.rawFrames.AddReplay()
	}
	return true
}

func ToggleReplaying() bool {
	if aiData.replaying == true {
		aiData.replaying = false
	}else{
		replayFrames := RawFramesToReplay(*aiData.rawFrames, 0, 0)
		*aiData.rawFrames = CBRRawFrames{}
		aiData.InitializeReplaying(replayFrames)
		aiData.replaying = true
	}
	return aiData.replaying
}

func CheckReplaying() bool{
	index := aiData.replayIndex
	ret := len(aiData.replayFrames) > index && aiData.replaying
	if ret == false {aiData.replaying = false}
	return ret
}

func ReadFrameInput(facing int32) int32 {
	index := aiData.replayIndex
	input := aiData.replayFrames[index].Input
	inp := input
	storedFacing := aiData.replayFrames[index].Facing
	bFacing := facingToBool(facing)
	if bFacing != storedFacing{
		input = swapBitsAtPos(input, 2, 3)
	}
	print("\n",facing, " ", bFacing, " ", storedFacing, " ", inp, " ", input)
	return input
}

func IncrementReplayIndex() bool {
	ret := false
	if aiData.replaying {
		aiData.replayIndex++
		ret = true
		print("\nFrame")
	}
	return ret
}

func AddFrameData(inputs []int32, facings []float32) bool {
	ret := false
	var bFacing bool
	if aiData.recording == true && len(aiData.rawFrames.ReplayFile) > 0 {
		aiData.rawFrames.AddFrame()
		aiData.rawFrames.AddCharData(len(inputs))
		for i := 0; i < len(inputs); i++ {
			bFacing = floatFacingToBool(facings[i])
			if i == 0 {print("\n", facings[0], " ",  inputs[0],  " ", bFacing)}
			aiData.rawFrames.SetPlayerInput(i, inputs[i], bFacing)

		}
		ret = true
	}
	return ret
}

func facingToBool(i int32) bool {
	var bFacing bool
	if i == -1 {
		bFacing = false
	} else if i == 1 {
		bFacing = true
	} else {
		print("There are more than 2 facings you dumbfuck go fix the function facingToBool")
	}
	return bFacing
}

func floatFacingToBool(f float32) bool {
	i := int(f)
	var bFacing bool
	if i == -1 {
		bFacing = false
	} else if i == 1 {
		bFacing = true
	} else {
		print("There are more than 2 facings you dumbfuck go fix the function facingToBool")
	}
	return bFacing
}

func boolToFacing(b bool) float32 {
	var iFacing float32
	if  b == false{
		iFacing = -1
	} else if  b == true {
		iFacing = 1
	} else {
		print("There are more than 2 facings you dumbfuck go fix the function boolToFacing")
	}
	return iFacing
}

///Interface functions for the outside program-----------

// switch left/right input for mirrored replay
func swapBitsAtPos(byte int32, posA int, posB int) int32{
	x := byte
	p1 := posA
	p2 := posB
	n := 1
	set1 :=  (x >> p1) & ((1 << n) - 1)
	set2 :=  (x >> p2) & ((1 << n) - 1)
	xor1 := set1 ^ set2
	xor := (xor1 << p1) | (xor1 << p2)
	result := x ^ xor
	return result
}

//Adds empty replay
func (x *CBRRawFrames) AddReplay() bool {
	//appending a new replay file to the CBR replay data everytime a new recording is started.
	x.ReplayFile = append(x.ReplayFile, &CBRRawFrames_ReplayFile{})
	return true
}
//Adds empty Frame
func (x *CBRRawFrames) AddFrame() bool {
	if x.ReplayFile != nil {
		var replayId = len(x.ReplayFile) - 1
		x.ReplayFile[replayId].Frame = append(x.ReplayFile[replayId].Frame, &CBRRawFrames_Frame{})
		//cbrData.ReplayFile[replayId].Frame[frameId].CharData = append(cbrData.ReplayFile[replayId].Frame[frameId].CharData, &CBRData_CharData{Input: input})
	}
	return true
}

//Adds empty CharData to Frame, every frame has data from multiple players stored
func (x *CBRRawFrames) AddCharData(players int) bool {
	if x.ReplayFile != nil {
		var replayId = len(x.ReplayFile) - 1
		var frameId = len(x.ReplayFile[replayId].Frame) - 1
		//cbrData.ReplayFile[replayId].Frame[frameId].CharData[playerNr].Input = input
		for i := 0; i < players; i++ {
			x.ReplayFile[replayId].Frame[frameId].CharData = append(x.ReplayFile[replayId].Frame[frameId].CharData, &CBRRawFrames_CharData{})
		}
	}
	return true
}

//Sets the data like player input and character facing in the the CharData of a frame
func (x *CBRRawFrames) SetPlayerInput(playerNr int, input int32, facing bool) bool {
	if x.ReplayFile != nil {
		var replayId = len(x.ReplayFile) - 1
		var frameId = len(x.ReplayFile[replayId].Frame) - 1
		if len(x.ReplayFile[replayId].Frame[frameId].CharData) > playerNr{
			x.ReplayFile[replayId].Frame[frameId].CharData[playerNr].Input = input
			x.ReplayFile[replayId].Frame[frameId].CharData[playerNr].Facing = facing
		}
	}
	return true
}

func (x *CBRRawFrames) clearRecording() bool {
	print("CBRRcord end")
	x.ReplayFile = nil
	return true
}

//Converts unprocessed frames to an array ready for replaying. Only player1 data is kept.
func RawFramesToReplay(frames CBRRawFrames, playerNr int, replayNr int) []*CBRData_Frame {
	var replayFrames []*CBRData_Frame
	for i := 0; i < len(frames.ReplayFile[replayNr].Frame); i++ {
		in := frames.ReplayFile[replayNr].Frame[i].CharData[playerNr].Input
		facing := frames.ReplayFile[replayNr].Frame[i].CharData[playerNr].Facing
		fr := CBRData_Frame{Input: in, Facing: facing}
		replayFrames = append(replayFrames, &fr)
	}
	return replayFrames
}

//Start replaying with given given array of frames
func (x *AIData) InitializeReplaying(frames []*CBRData_Frame) bool {
	ret := false
	if frames != nil && len(frames) > 0 {
		x.replayFrames = frames
		x.replaying = true
		x.replayIndex = 0
		ret = true
	}
	return ret
}


//Starts recording in a new replay File
func (x *AIData) InitializeRecording() bool {
	print("CBRRcord start")
	if x.rawFrames.ReplayFile == nil{
		x.rawFrames = &CBRRawFrames{ReplayFile: []*CBRRawFrames_ReplayFile{}}
	}
	x.recording = true
	return true
}

func (x *AIData) EndRecording() bool {
	print("CBRRcord end")
	x.recording = false
	return true
}


func (x *CBRData) Initalize() bool {
	//sys.cbrData = cbrData
	//CBRData.ReplayFile := cbr.CBRData_ReplayFile{}
	//var replayId = len(sys.cbrData.ReplayFile) - 1
	//sys.cbrData.ReplayFile[replayId].Frame = append(sys.cbrData.ReplayFile[0].Frame, &cbr.CBRData_Frame{})
	//Creating or loading a cbrData structure when starting to record for the AI and saving the recorded data when stopping the recording
	//var data  []byte
	//var err error
	if x == nil{
		print("CBRRcord start")
		data, err := ioutil.ReadFile("CBRReplays.data")
		if err != nil {
			x = &CBRData{
				ReplayFile: []*CBRData_ReplayFile{},
			}
		}else{
			x = &CBRData{
				ReplayFile: []*CBRData_ReplayFile{},
			}
			err := proto.Unmarshal(data, x)
			if err != nil {
				log.Fatal("unmarshaling error: ", err)
			}
		}
	}else{
		x.CBRClose()
	}
	return true
}

func (x *CBRData) CBRClose() bool {
	print("CBRRcord end")
	data, err := proto.Marshal(x)
	err = proto.Unmarshal(data, x)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	ioutil.WriteFile("CBRReplays.data", data, 0644)
	x = nil
	return true
}

