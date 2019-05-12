package gospeak

import (speech https://github.com/sveggen/go-speak)
//example


func main(){
  speech.SetWitKey("RLD2DHGWIRNX72FHOZMRNBTGJBGOCEXK") //Wit API key fra sveggen
  speech.SendWitVoice("Jorgen.flac")
  

}


