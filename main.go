package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/mail"
	"strings"
)
func BodyParser( r *http.Request, x interface{}){

if body, err:= io.ReadAll(r.Body) ; err==nil{
  if err:=json.Unmarshal([]byte(body), x); err!=nil{
	return
  }
}


}
func Validator(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path!="/verifyEmail" {
         http.Error(w,"Incorrect Endpoint", http.StatusNotFound)
		 return
		}
      if r.Method!=http.MethodPost{
		http.Error(w,"Method not allow", http.StatusMethodNotAllowed)
		return
	  }
   

	  next(w,r)
	}
}

type ResponseStruct struct{
	Email string `json:"email"`
	MX bool `json:"mx"`
	DMARC bool `json:"dmarc"`
	SPF bool `json:"spf"`
	SPFRecord string `json:"SPF_Record"`
	DMARCRecord string `json:"DMARC_Record"`
	CatchAllSetUp bool `json:"Catch_All_Set_up"`
    RoleBased bool `json:"Role_Based"`
	Role string `json:"Role"`
	SMTP bool `json:"SMTP Verfied"`
	WebsiteAssociation bool `json:"website_Association"`
	Websites string `json:"Websites"`
}
var resobj ResponseStruct
var DisposableEmailProviders = []string{
	"mailinator.com",
	"guerrillamail.com",
	"dispostable.com",
	"10minutemail.com",
	"temp-mail.org",
	"maildrop.cc",
	"tempmailaddress.com",
	"mailnesia.com",
	"my10minutemail.com",
	"throwawaymail.com",
	"sharklasers.com",
	"mailinator2.com",
	"yopmail.com",
	"temp-mail.ru",
	"mailinator.net",
	"mailcatch.com",
	"mailmetrash.com",
	"fakeinbox.com",
	"pookmail.com",
	"discard.email",
	"mailnull.com",
	"meltmail.com",
	"zomg.info",
	"jetable.org",
	"trashmail.com",
	"spamgourmet.com",
	"mintemail.com",
	"mailinator.co.uk",
	"mailinator.ca",
	"mailinator.info",
	"mailinator.com",
}
var RoleEmailPrefixes = []string{"info", "support", "sales", "admin"}
func IsDisposableEmail(email string)bool{
 part:=strings.Split(email, "@")
 domain:=part[1]

 for _, provider:=range DisposableEmailProviders{
	if strings.EqualFold(domain,provider){
		return true
	}
 }
 return false
}

func ExtractDomain(email string) (string,string) {
   
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
       
        return "",""
    }
   
    return parts[0],parts[1]
}

func SyntaxCheck(email string) bool{
 _,err:=mail.ParseAddress(email)
 return err==nil
}

func WebsiteExists(domain string) ([]string, error) {
  
    ips, err := net.LookupIP(domain)
    if err != nil {
        return nil, err
    }
    var websiteNames []string
    for _, ip := range ips {
		
		 if err!=nil{
			fmt.Println(err,"err of host")
		 }
		 websiteNames=append(websiteNames, ip.String())
     
     }
    return websiteNames, nil
}

func EmailVerifier(email string, resobj ResponseStruct) ResponseStruct{
    resobj.Email=email

	if (len(resobj.Email)>0){
     disposable:=IsDisposableEmail(resobj.Email)
     correctEmail:=SyntaxCheck(resobj.Email)

	if correctEmail && !disposable {
		fmt.Print(correctEmail, disposable)
    // role,domain:=ExtractDomain(resobj.Email)
	// mxRecord,err:=net.LookupMX(domain)
	// if err!=nil{
	// 	fmt.Print("Error comming 59")
		
	// }
	// if len(mxRecord) == 0{
    //  resobj.CatchAllSetUp=true
	// }
	// if(len(mxRecord)>0){
	// 	resobj.MX=true
	// }

	// websites, err := WebsiteExists(domain)
	// if err == nil {
		
	// 	resobj.WebsiteAssociation=true
	// 	resobj.Websites=strings.Join(websites,",")
	// }
	
	
    //  smtpServer := mxRecord[0].Host
    //  smtpPort := "25" 
	//  fmt.Print(smtpServer,"smtpServer \n",smtpPort,"\n smtpPort")
	 
    // client, err := smtp.Dial(smtpServer+":"+smtpPort)
    // if err != nil {
    //     fmt.Println("Failed to connect to SMTP server:", err)
    //     return
    // }
    // defer client.Close()
	
     
	//   if err==nil{
	// 	fmt.Print("Verification in process")
	// 	client.Hello(resobj.Email)
	// 	err:=client.Mail(resobj.Email)
	// 	if err==nil{
	// 		err = client.Rcpt(resobj.Email)
	// 		if err==nil{
	// 			resobj.SMTP=true
	// 		}
	// 	}
		
	//  }
   
	//  txtRecord,err:=net.LookupTXT(domain)
	// if err!=nil{
	// 	fmt.Print("Error comming 59")
		
	// }
	// for _,record:=range txtRecord{
	// 	if strings.HasPrefix(record,"v=spf1"){
	// 		resobj.SPF=true
	// 		resobj.SPFRecord=record
	// 		break
	// 	}

	// }

	// dmarkRecord,err:=net.LookupTXT("_dmarc."+domain)

	// if err!=nil{
	// 	fmt.Print("Error comming 94")		
	// }
	// for _,record:=range dmarkRecord{
	// 	if strings.HasPrefix(record,"v=DMARC1"){
	// 		resobj.DMARC=true
	// 		resobj.DMARCRecord=record
	// 		break
	// 	}

	// }

	// for _, pref:=range RoleEmailPrefixes{
	// 	if strings.EqualFold(role, pref){
	// 		resobj.RoleBased=true
	// 		resobj.Role=role
	// 		break;
	// 	}
	// }
	 }
	}
	

	return resobj
}
func SingleEmailVerifier(w http.ResponseWriter, r *http.Request, resobj *ResponseStruct){
   
	BodyParser(r,&resobj)
	fmt.Print(resobj,"mai set ho gya hu")
	// if !(len(resobj.Email)>0){
    //   http.Error(w,"Bad Request",http.StatusBadRequest)
	//   return
	// }
	disposable:=IsDisposableEmail(resobj.Email)
    correctEmail:=SyntaxCheck(resobj.Email)

	if correctEmail && !disposable {
    role,domain:=ExtractDomain(resobj.Email)
	mxRecord,err:=net.LookupMX(domain)
	if err!=nil{
		fmt.Print("Error comming 59")
		
	}
	if len(mxRecord) == 0{
     resobj.CatchAllSetUp=true
	}
	if(len(mxRecord)>0){
		resobj.MX=true
	}

	websites, err := WebsiteExists(domain)
	if err == nil {
		
		resobj.WebsiteAssociation=true
		resobj.Websites=strings.Join(websites,",")
	}
	
	
     smtpServer := mxRecord[0].Host
     smtpPort := "25" 
	 fmt.Print(smtpServer,"smtpServer \n",smtpPort,"\n smtpPort")
	 
    // client, err := smtp.Dial(smtpServer+":"+smtpPort)
    // if err != nil {
    //     fmt.Println("Failed to connect to SMTP server:", err)
    //     return
    // }
    // defer client.Close()
	
     
	//   if err==nil{
	// 	fmt.Print("Verification in process")
	// 	client.Hello(resobj.Email)
	// 	err:=client.Mail(resobj.Email)
	// 	if err==nil{
	// 		err = client.Rcpt(resobj.Email)
	// 		if err==nil{
	// 			resobj.SMTP=true
	// 		}
	// 	}
		
	//  }
   
	 txtRecord,err:=net.LookupTXT(domain)
	if err!=nil{
		fmt.Print("Error comming 59")
		
	}
	for _,record:=range txtRecord{
		if strings.HasPrefix(record,"v=spf1"){
			resobj.SPF=true
			resobj.SPFRecord=record
			break
		}

	}

	dmarkRecord,err:=net.LookupTXT("_dmarc."+domain)

	if err!=nil{
		fmt.Print("Error comming 94")		
	}
	for _,record:=range dmarkRecord{
		if strings.HasPrefix(record,"v=DMARC1"){
			resobj.DMARC=true
			resobj.DMARCRecord=record
			break
		}

	}

	for _, pref:=range RoleEmailPrefixes{
		if strings.EqualFold(role, pref){
			resobj.RoleBased=true
			resobj.Role=role
			break;
		}
	}
	}
	 
}

func handleExcelUpload(w http.ResponseWriter, r *http.Request){
  file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	var emails []ResponseStruct

	// Read each line of the file
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}
		lines := strings.Split(string(buf[:n]), "\n")
		for _, line := range lines {
			// Perform email verification
			email := strings.TrimSpace(line)
			var obj ResponseStruct
			 EmailVerifier(email, obj)
			// fmt.Println(resobj,"resobj")
			// Add the email and verification status to the slice
			emails = append(emails, obj)
		}
	}

	// Marshal the array of EmailVerification structs to JSON
	responseData, err := json.Marshal(emails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func EmailHandler(w http.ResponseWriter, r *http.Request){
	
   // limiting to file upload to approx 10 mb
	// err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	// if err == nil && r.MultipartForm != nil && r.MultipartForm.File != nil {
		
	// 	handleExcelUpload(w, r)
	// 	return
	// }


	SingleEmailVerifier(w,r,&resobj)

	res,_ :=json.Marshal(resobj)
    w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
    w.Write(res)
	
}

func main() {


	 http.Handle("/verifyEmail",Validator(EmailHandler))

     fmt.Print("Server is running on port 8080")
	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal("Error in listning\n", err)
	}else{
		fmt.Print("	Server is running on port 8080")
	}

}