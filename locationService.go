package main

import (
	
	"encoding/json"
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"	
    "io/ioutil"
    s "strings"
    "math/rand"
    "strconv"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "time"
)

type ArgsForCreatingLocation struct {
    Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
}



type Response struct{
	
	Id string `json:"id"`
	Name string `json:"name"`
    Address string `json:"address"`
    City string `json:"city"`
    State string `json:"state"`
    Zip string `json:"zip"`
    Coordinate Coor`json:"coordinate"`
  }
  type Coor struct {
          Lat string `json:"lat"`
          Lng string `json:"lng"`
      }
       
   

func Createlocations(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
        
    //var resp Response
    var request ArgsForCreatingLocation
    var Addr string
    var city string
    var state string
    var zip string
    json.NewDecoder(req.Body).Decode(&request)
    
    if(s.Contains(request.Address," ")){
    	 Addr = s.Replace(request.Address," ","+",-1)	
    	}
    if(s.Contains(request.City," ")){
    	 city = s.Replace(request.City," ","+",-1)	
    	}
    if(s.Contains(request.State," ")){
    	 state = s.Replace(request.State," ","+",-1)	
    	}
    if(s.Contains(request.Zip," ")){
    	 zip = s.Replace(request.Zip," ","+",-1)	
    	}
    
	url := "http://maps.google.com/maps/api/geocode/json?address="+Addr+city+state+zip
    resp,_:= http.Get(url)
    //var reply Response
    defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
    var f interface{}

    error := json.Unmarshal(body,&f)
    fmt.Println(error)
   	


    mRes := f.(map[string]interface{})["results"]
    mRes0 := mRes.([]interface{})[0]
    mGeo := mRes0.(map[string]interface{})["geometry"]
    mLoc := mGeo.(map[string]interface{})["location"]

    locLat := mLoc.(map[string]interface{})["lat"].(float64)
    locLng := mLoc.(map[string]interface{})["lng"].(float64) 
	
	mongoDBDialInfo := &mgo.DialInfo{
	Addrs:    []string{"ds043694.mongolab.com:43694"},
	Timeout:  60 * time.Second,
	Database: "locations",
	Username: "Admin",
	Password: "Admin123",
}    

    session, err := mgo.DialWithInfo(mongoDBDialInfo)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
        c := session.DB("locations").C("LocDetails")
        

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
	var response Response
	lat := strconv.FormatFloat(locLat, 'f', 6, 64)
	lng := strconv.FormatFloat(locLng, 'f', 6, 64)
	newID := strconv.Itoa(rand.Intn(99999))
	response.Id = newID
	response.Name = request.Name 
	response.Address = request.Address
	response.City = request.City
	response.State=request.State
	response.Zip = request.Zip
	response.Coordinate.Lat = lat
	response.Coordinate.Lng = lng
	err = c.Insert(response)
        if err != nil {
                log.Fatal(err)
        }


   
    uj, _ := json.Marshal(response)
    fmt.Fprintf(w, "%s",uj)
}

func GetLocation(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	IdToGet := p.ByName("id")
	result := Response{}
	mongoDBDialInfo := &mgo.DialInfo{
	Addrs:    []string{"ds043694.mongolab.com:43694"},
	Timeout:  60 * time.Second,
	Database: "locations",
	Username: "Admin",
	Password: "Admin123",
}    

    session, err := mgo.DialWithInfo(mongoDBDialInfo)
    if err != nil {
        panic(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)
    c := session.DB("locations").C("LocDetails")

	err = c.Find(bson.M{"id": IdToGet}).One(&result)
	if err != nil {
		panic(err)
	}
	var response Response
	response.Id = result.Id
	response.Name = result.Name 
	response.Address = result.Address
	response.City = result.City
	response.State=result.State
	response.Zip = result.Zip
	response.Coordinate.Lat = result.Coordinate.Lat
	response.Coordinate.Lng = result.Coordinate.Lng
	uj, _ := json.Marshal(response)
    fmt.Fprintf(w, "%s",uj)

	
}
func UpdateLocation(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var request ArgsForCreatingLocation
	Data := Response{}
    var Addr string
    var city string
    var state string
    var zip string
    json.NewDecoder(req.Body).Decode(&request)
    
    if(s.Contains(request.Address," ")){
    	 Addr = s.Replace(request.Address," ","+",-1)	
    	}
    if(s.Contains(request.City," ")){
    	 city = s.Replace(request.City," ","+",-1)	
    	}
    if(s.Contains(request.State," ")){
    	 state = s.Replace(request.State," ","+",-1)	
    	}
    if(s.Contains(request.Zip," ")){
    	 zip = s.Replace(request.Zip," ","+",-1)	
    	}
    IdToGet := p.ByName("id")

    url := "http://maps.google.com/maps/api/geocode/json?address="+Addr+city+state+zip
    resp,_:= http.Get(url)
    //var reply Response
    defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
    var f interface{}

    error := json.Unmarshal(body,&f)
    fmt.Println(error)
   	


    mRes := f.(map[string]interface{})["results"]
    mRes0 := mRes.([]interface{})[0]
    mGeo := mRes0.(map[string]interface{})["geometry"]
    mLoc := mGeo.(map[string]interface{})["location"]

    locLat := mLoc.(map[string]interface{})["lat"].(float64)
    locLng := mLoc.(map[string]interface{})["lng"].(float64) 
	
	mongoDBDialInfo := &mgo.DialInfo{
	Addrs:    []string{"ds043694.mongolab.com:43694"},
	Timeout:  60 * time.Second,
	Database: "locations",
	Username: "Admin",
	Password: "Admin123",
}    

    session, err := mgo.DialWithInfo(mongoDBDialInfo)
        if err != nil {
                panic(err)
        }
        defer session.Close()

    session.SetMode(mgo.Monotonic, true)
    c := session.DB("locations").C("LocDetails")
        

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
	var response Response
	lat := strconv.FormatFloat(locLat, 'f', 6, 64)
	lng := strconv.FormatFloat(locLng, 'f', 6, 64)
	err = c.Find(bson.M{"id": IdToGet}).One(&Data)
	if err != nil {
		panic(err)
	}
	
	response.Id = IdToGet
	response.Name = Data.Name 
	response.Address = request.Address
	response.City = request.City
	response.State=request.State
	response.Zip = request.Zip
	response.Coordinate.Lat = lat
	response.Coordinate.Lng = lng
	
  
    err = c.Update(Data, response)
	if err != nil {
		panic(err)
	}
   
    uj, _ := json.Marshal(response)
    fmt.Fprintf(w, "%s",uj)
	
}
func DeleteLocation(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	IdToGet := p.ByName("id")
	result := Response{}
	mongoDBDialInfo := &mgo.DialInfo{
	Addrs:    []string{"ds043694.mongolab.com:43694"},
	Timeout:  60 * time.Second,
	Database: "locations",
	Username: "Admin",
	Password: "Admin123",
}    

    session, err := mgo.DialWithInfo(mongoDBDialInfo)
    if err != nil {
        panic(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)
    c := session.DB("locations").C("LocDetails")

	err = c.Find(bson.M{"id": IdToGet}).One(&result)
	if err != nil {
		panic(err)
	}
	err = c.Remove(result)
	if err != nil {
        panic(err)
    }
    
}

func main() {
    mux := httprouter.New()
    mux.POST("/locations", Createlocations)
    mux.GET("/locations/:id",GetLocation)
    mux.PUT("/locations/:id",UpdateLocation)
    mux.DELETE("/locations/:id",DeleteLocation)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}

