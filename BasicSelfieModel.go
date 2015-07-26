package main
import "fmt"
 
type Instagram struct {
    barzoCount 	int
	likeCount 	int
	gender 		string
	comments 	[]string
	isFamous 	bool
}

func NewInstagram(gender string) Instagram{
	this := Instagram{}
	this.barzoCount = 50
	this.isFamous = true
	this.likeCount = 1
	this.gender = gender
	return this
}
 
type Selfie struct{
	Instagram
	showBoobs bool
}

func (this *Selfie) takePhoto(showBoobs bool){
	this.showBoobs = showBoobs
	
	if showBoobs {
		this.isFamous = true
	}
	
	if this.gender == "male" && !this.isFamous {
		this.likeCount = 7
	} else if this.gender == "male" && this.isFamous {
		this.likeCount = 300 - this.barzoCount
	} else if this.gender == "famale" {
		if !this.isFamous {
			this.likeCount = 20
		} else {
			this.likeCount = 700
		}
		
		for this.showBoobs && this.barzoCount < 150 {
			this.comments = append(this.comments,"Memeni Aç böğğhhğh")
			this.likeCount += this.barzoCount
			this.barzoCount++
		}
		
	} else {
		fmt.Println("aq faşisti eray az pride pls")
	}
}

func main() {
	Pelin := Selfie{NewInstagram("famale"),true}
	Pelin.takePhoto(true)
}
