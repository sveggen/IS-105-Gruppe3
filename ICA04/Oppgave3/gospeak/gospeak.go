package main

import(speech "./go-speak")


func main(){
  speech.SetWitKey("RLD2DHGWIRNX72FHOZMRNBTGJBGOCEXK") //Wit API key fra sveggen
  speech.SendWitVoice("Jorgen.wav")
  

}


