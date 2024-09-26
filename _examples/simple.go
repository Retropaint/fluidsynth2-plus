package main

import (
	"fmt"

	"github.com/retropaint/fluidsynth2-plus"
)

func main() {

	s := fluidsynth2.NewSettings()
	fmt.Println("Avaliable audio drivers:")
	for _, driver := range s.GetOptions("audio.driver") {
		fmt.Print(driver + " ")
	}
	// Easy way to set audio backend
	//s.SetString("audio.driver", "portaudio")

	synth := fluidsynth2.NewSynth(s)

	synth.SFLoad("Roland_SC-88.sf2", true) // https://musical-artifacts.com/artifacts/538

	// Extra synth options
	// (Reverb and chorus disabled)
	synth.SetReverb(0, 0)
	synth.SetChorus(0, 0)

	player := fluidsynth2.NewPlayer(synth)
	player.Add("sample.mid")

	// Example of how to play from memory
	// dat, err := ioutil.ReadFile("midifile.mid")
	// if err != nil {
	// 	panic(err)
	// }

	// player.AddMem(dat)

	// Easy way to set audio backend
	//s.SetString("audio.driver", "coreaudio")

	fluidsynth2.NewAudioDriver(s, synth)

	player.Play()

	// Increase tempo
	// player.SetTempo(fluidsynth2.TEMPO_INTERNAL, 2)

	player.Join()

}
